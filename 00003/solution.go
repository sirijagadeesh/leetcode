package leedcode

import (
	"strings"
)

func lengthOfLongestSubstring(s string) int {
	current, largest := "", ""

	for _, v := range s {
		if i := strings.IndexRune(current, v); i >= 0 {
			if len(current) > len(largest) {
				largest = current
			}
			current = current[i+1:]
		}
		current += string(v)
	}

	if len(current) > len(largest) {
		return len(current)
	}

	return len(largest)
}
