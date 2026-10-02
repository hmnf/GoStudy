package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(Search(
		[]int{5},
		5,
	))
}

func Search(nums []int, target int) int {
	if len(nums) == 1 {
		if nums[0] == target {
			return 0
		} else {
			return -1
		}
	}
	left := 0
	right := len(nums) - 1

	for right >= left {
		mid := int(math.Floor(float64(right+left) / 2))
		fmt.Println(mid)
		if nums[mid] == target {
			return mid
		} else if nums[mid] > target {
			right = mid
		} else {
			left = mid
		}
	}

	return -1
}
