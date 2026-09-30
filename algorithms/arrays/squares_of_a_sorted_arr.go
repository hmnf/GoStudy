package main

func SquaresOfASortedArrays(nums []int) []int {
	left := 0
	right := len(nums) - 1
	pointer := len(nums) - 1

	res := make([]int, len(nums))

	for right > left {
		if nums[left]*nums[left] > nums[right]*nums[right] {
			res[pointer] = nums[left] * nums[left]
			left++
		} else {
			res[pointer] = nums[right] * nums[right]
			right--
		}

		pointer--
	}

	return res
}
