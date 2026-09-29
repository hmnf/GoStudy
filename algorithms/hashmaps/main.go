package main

import "fmt"

func main() {
	fmt.Println(
		TwoSum([]int{2, 7, 11, 15}, 9),
		TwoSum([]int{3, 3}, 6),
		TwoSum([]int{3, 2, 4}, 6),
		isAnagram("anagram", "nagaram"),
	)
}

func isAnagram(s string, t string) bool {
	count := make(map[rune]int)

	for _, l := range s {
		count[l]++
	}

	for _, l := range t {
		count[l]--
		if count[l] == 0 {
			delete(count, l)
		}
	}

	if len(count) == 0 {
		return true
	}

	return false
}
