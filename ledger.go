package fixture

import "errors"

type Transfer struct { ID, From, To string; Amount int }
type Ledger struct { Balances map[string]int; Seen map[string]bool }

func (l *Ledger) ApplyBatch(batch []Transfer) error {
    if l.Balances == nil { l.Balances = map[string]int{} }
    if l.Seen == nil { l.Seen = map[string]bool{} }

    // Validate the batch first for atomicity
    seenIDs := make(map[string]bool)
    tempBalances := make(map[string]int)
    for k, v := range l.Balances {
        tempBalances[k] = v
    }

    for _, t := range batch {
        if t.ID == "" || t.From == "" || t.To == "" || t.Amount <= 0 || t.From == t.To { return errors.New("invalid transfer") }
        if l.Seen[t.ID] || seenIDs[t.ID] { return errors.New("duplicate id") }
        seenIDs[t.ID] = true

        if tempBalances[t.From] < t.Amount { return errors.New("insufficient funds") }
        tempBalances[t.From] -= t.Amount
        tempBalances[t.To] += t.Amount
    }

    // Apply changes
    for _, t := range batch {
        l.Balances[t.From] -= t.Amount
        l.Balances[t.To] += t.Amount
        l.Seen[t.ID] = true
    }
    return nil
}
