package leetcode

import (
	"fmt"
	"regexp"
)

func isMatch(s string, p string) bool {
	ok, err := regexp.MatchString(fmt.Sprintf("^%s$", p), s)
	if err != nil {
		return false
	}

	return ok
}
