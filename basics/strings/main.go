package main

import "fmt"

func main() {
	fmt.Println(
		Reverse("hello"),
		Reverse("go"),
		Reverse(""),
		Reverse("привет"),
		FirstUnique("leetcode"),
		FirstUnique("loveleetcode"),
	)
}

func Reverse(s string) string {
	r := make([]rune, len(s))
	i := len(s)
	for _, l := range []rune(s) {
		r[i-1] = l
		i--
	}

	return string(r)
}
