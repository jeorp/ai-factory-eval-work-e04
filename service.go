package fixture
import "strings"
func Add(a,b int) int {return a+b}
func Valid(s string) bool {return strings.TrimSpace(s) != ""}
func Normalize(s string) string {return strings.ToLower(strings.TrimSpace(s))}
func Lookup(m map[string]string,k string) string {if m == nil { return "" }; val, ok := m[k]; if !ok { return "" }; return val }
func Version() string {return "v1"}
var _ = strings.TrimSpace
