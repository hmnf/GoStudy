package main

import "fmt"

func main() {
	fmt.Println(
		frequency([]int{1, 2, 2, 3, 3, 3}),
		unique([]int{1, 2, 2, 3, 1}),
	)
}

func frequency(nums []int) map[int]int {
	freq := make(map[int]int)

	for _, num := range nums {
		freq[num]++
	}

	return freq
}

func unique(nums []int) []int {
	set := make(map[int]struct{})

	for _, num := range nums {
		set[num] = struct{}{}
	}

	var uniq []int

	for v := range set {
		uniq = append(uniq, v)
	}

	return uniq
}
