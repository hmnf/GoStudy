package main

import "fmt"

func main() {
	nums := []int{0, 1, 0, 3, 12}

	moveZeroes(nums)

	fmt.Println(nums)
}

func moveZeroes(nums []int) {
	right := len(nums) - 1
	left := right - 1

	for left >= 0 {
		if nums[left] == 0 {
			for i := left; i < right; i++ {
				nums[i], nums[i+1] = nums[i+1], nums[i]
			}
		}
		left--
	}
}
