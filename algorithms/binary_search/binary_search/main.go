package main

import (
	"fmt"
)

func main() {
	fmt.Println(Search(
		[]int{5},
		5,
	))
}

func Search(nums []int, target int) int {
	left := 0
	right := len(nums) - 1

	for right >= left {
		mid := (right + left) / 2
		if nums[mid] == target {
			return mid
		} else if nums[mid] > target {
			right = mid - 1
		} else {
			left = mid + 1
		}
	}

	return -1
}
