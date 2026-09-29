package main

func IsPalindrome(s string) bool {
	rn := []rune(s)
	l := 0
	r := len(rn) - 1

	for r > l {
		if rn[r] != rn[l] {
			return false
		}
		r--
		l++
	}

	return true
}
