package main

func IsPalindrome(s string) bool {
	l := 0
	r := len(s) - 1

	for r > l {
		if s[r] != s[l] {
			return false
		}
		r--
		l++
	}

	return true
}
