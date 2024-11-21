// checklist/identity/cmd.go
package identity

import (
	"fmt"
	"strings"

	"github.com/gunh0/openstack-security-hub/checklist"
)

func prettyPrintResult(result checklist.CheckResult) {
	fmt.Println(strings.Repeat("-", 100))
	fmt.Printf("Description: %s\n", result.Description)
	fmt.Printf("Result: %s\n", result.Result)
	fmt.Printf("Details: %s\n", result.Details)
	fmt.Println(strings.Repeat("-", 100))
}
