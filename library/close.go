package library

import (
	"fmt"
	"os"
	"strings"
)

func close(format string, k ...interface{}) bool {
	if format != "" {
		n := `
`
		if !strings.Contains(format, n) {
			format = fmt.Sprintf("%s%s", format, n)
		}
		fmt.Printf(format, k...)
		os.Exit(1)
		return false
	}
	return true
}
