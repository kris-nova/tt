package library

import (
	"fmt"
)

func init() {

	// -----------------------------------------------------------------
	// StdOut
	stdLib["StdOut"] = func(q string) bool {
		str, err := lookup(q)
		if err != nil {
			return close(err.Error())
		}
		fmt.Println(str)
		return close("")
	}
}
