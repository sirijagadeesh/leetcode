package leetcode

func isPalindrome(x int) bool {
	if x < 0 {
		return false
	}

	return x == reverseNumber(x)
}

func reverseNumber(x int) int {
	y := 0

	for x != 0 {
		l := x % 10
		y = y*10 + l
		x /= 10
	}

	return y
}
