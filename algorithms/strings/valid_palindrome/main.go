package main

import "fmt"

func main() {
	s := "A man, a plan, a canal: Panama"
	fmt.Println(ValidPalindrome(s))
}

func ValidPalindrome(s string) bool {
	var b []rune

	for _, r := range s {
		if r >= 65 && r <= 90 {
			b = append(b, r+32)
		} else if r >= 97 && r <= 122 {
			b = append(b, r)
		}
	}

	left := 0
	right := len(b) - 1

	for right >= left {
		if b[right] != b[left] {
			return false
		}
		right--
		left++
	}

	return true
}
