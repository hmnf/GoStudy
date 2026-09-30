package main

func SortColors(nums []int) {
	left := 0
	right := len(nums) - 1
	current := left

	for current <= right {
		switch nums[current] {
		case 0:
			nums[left], nums[current] = nums[current], nums[left]
			left++
			current++
		case 2:
			nums[right], nums[current] = nums[current], nums[right]
			right--
		default:
			current++
		}
	}
}
