package stackerr

import "strings"

func Is(in error, values ...error) bool {
	cause := Cause(in)
	root := Root(in)

	for _, v := range values {
		if cause == v || root == v {
			return true
		}
	}

	errStr := in.Error()

	for _, v := range values {
		if strings.Contains(errStr, v.Error()) {
			return true
		}
	}

	return false
}
