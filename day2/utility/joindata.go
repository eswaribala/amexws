package utility //reusable package
import (
	"strings"
)

func ToString(customerData []string) string {

	return strings.Join(customerData, "->")

}
