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
	res := make(map[rune]int)

	for _, l := range s {
		res[l]++
	}

	for _, l := range t {
		res[l]--
		if res[l] == 0 {
			delete(res, l)
		}
	}

	if len(res) == 0 {
		return true
	}

	return false
}
