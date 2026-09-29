package main

import (
	"fmt"
)

func MergeSortedArrays(nums1 []int, m int, nums2 []int, n int) []int {
	var i, j int

	if m > n {
		i = m - 1
		j = n - 1
	}
	i = n - 1
	j = m - 1
	nums1, nums2 = nums2, nums1
	for j >= 0 && i >= 0 {
		if nums2[j] >= nums1[i] {
			nums1 = append(nums1, 0)
			fmt.Println(nums1, nums2, i, j)
			copy(nums1[i+2:], nums1[i+1:])
			nums1[i+1] = nums2[j]
			nums2 = nums2[:len(nums2)-1]
			fmt.Println(nums2)
			j--
		} else {
			i--
		}
	}

	if j >= 0 {
		return append(nums2, nums1...)
	}

	return nums1
}
