package main

import "fmt"

func main() {
	nums := []int{5, 2, 4, 6, 1, 3}

	insertionSort(nums)

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
