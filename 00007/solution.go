package leetcode

import (
	"math"
)

func reverse(x int) int {
	reversed := 0
	reminder := 0
	negitive := false

	if x < 0 {
		negitive = true
		x *= -1
	}

	for x > 0 {
		reminder = x % 10
		reversed = reversed*10 + reminder
		x /= 10
	}

	if reversed > math.MaxInt32 || reversed < math.MinInt32 {
		return 0
	}

	if negitive {
		return reversed * -1
	}

	return reversed
}
