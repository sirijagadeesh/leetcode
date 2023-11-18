package leetcode

import (
	"math"
	"regexp"
	"strconv"
	"strings"
)

func myAtoi(s string) int {
	re := regexp.MustCompile(`^[-+]?[0-9]+`)
	s = strings.TrimSpace(s)
	matches := re.FindStringSubmatch(s)

	if len(matches) > 0 {
		num, _ := strconv.Atoi(matches[0])

		if num < math.MinInt32 {
			return math.MinInt32
		}

		if num > math.MaxInt32 {
			return math.MaxInt32
		}

		return num
	}

	return 0
}
