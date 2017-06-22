package library

import (
	"fmt"
	"strings"
)

var vars = make(map[string]string)

func init() {

	// -----------------------------------------------------------------
	// SetVar
	stdLib["SetVar"] = func(key string, q string) bool {
		str, err := lookup(q)
		if err != nil {
			return close(err.Error())
		}
		vars[key] = str
		return close("")
	}

	// -----------------------------------------------------------------
	// VarDump
	stdLib["VarDump"] = func() bool {
		for k, v := range vars {
			fmt.Printf("%s: %s\n", k, v)
		}
		return close("")
	}
}

func lookup(q string) (string, error) {
	if strings.HasPrefix(q, "$") {
		nq := strings.TrimPrefix(q, "$")
		return lookup(nq)
	} else {
		if val, ok := vars[q]; ok {
			return lookup(val)
		} else {
			return q, nil
		}
	}
	return "", nil
}

func recursiveLookup() {

}
