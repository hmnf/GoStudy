package main

func insertionSort(nums []int) {
	for i := 1; i < len(nums); i++ {
		key := nums[i]
		for j := i - 1; j >= 0; j-- {
			if key < nums[j] {
				nums[j+1] = nums[j]
				nums[j] = key
			} else {
				break
			}
		}
	}
}
