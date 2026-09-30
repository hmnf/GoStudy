package main

func insertionSort(nums []int) {
	for i := 0; i < len(nums); i++ {
		for j := i - 1; j >= 0; j-- {
			if nums[i] < nums[j] {
				nums[i], nums[j] = nums[j], nums[i]
				i--
			}
		}
	}
}
