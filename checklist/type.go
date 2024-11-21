// checklist/type.go
package checklist

// CheckResult represents a check result
type CheckResult struct {
	Description string `json:"description"`
	Result      string `json:"result"`
	Details     string `json:"details"`
}
