package main

import "fmt"

func main() {
	nums := []int{2, 0, 2, 1, 1, 0}

	SortColors(nums)

	fmt.Println(nums)
}

func arrEqual(got, want []int) bool {
	if len(got) != len(want) {
		return false
	}

	for i := 0; i < len(got); i++ {
		if got[i] != want[i] {
			return false
		}
	}

	return true
}
