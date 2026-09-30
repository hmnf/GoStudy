package main

import "fmt"

func SortColors(nums []int) {
	left := 0
	right := len(nums) - 1
	for nums[right] == 2 && right > 0 {
		right--
	}
	for nums[left] == 0 && left < len(nums)-1 {
		left++
	}
	for nums[right] == 0 && left < len(nums)-1 {
		nums[right], nums[left] = nums[left], nums[right]
		if nums[right] == 2 {
			right--
		}
		left++
	}

	current := left

	for current <= right {
		fmt.Println(left, right)
		switch nums[current] {

		case 0:
			nums[left], nums[current] = nums[current], nums[left]
			left++
		case 2:
			nums[right], nums[current] = nums[current], nums[right]
			right--
		}
		fmt.Println(nums, nums[current])
		current++
	}
}
