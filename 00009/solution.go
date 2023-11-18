package leetcode

func isPalindrome(x int) bool {
	z := x
	if z < 0 {
		return false
	}

	y := 0

	for z != 0 {
		l := z % 10
		y = y*10 + l
		z /= 10
	}

	return x == y
}
