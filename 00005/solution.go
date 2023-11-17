package leetcode

func longestPalindrome(s string) string {
	res := ""

	for i := range s {
		res = maxPalindrome(s, i, i, res)   // Odd length palindrome
		res = maxPalindrome(s, i, i+1, res) // Even length palindrome
	}

	return res
}

func maxPalindrome(s string, left, right int, cur string) string {
	newCurr := ""

	for ; left >= 0 && right < len(s) && s[left] == s[right]; left, right = left-1, right+1 {
		newCurr = s[left : right+1]
	}

	if len(cur) < len(newCurr) {
		return newCurr
	}

	return cur
}
