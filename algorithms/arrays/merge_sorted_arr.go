package main

import "fmt"

func MergeSortedArrays(nums1 []int, m int, nums2 []int, n int) []int {
	k := len(nums1) - 1
	n--
	m--
	fmt.Println(k, m, n)
	for k >= 0 && m >= 0 && n >= 0 {
		if nums2[n] > nums1[m] {
			nums1[k] = nums2[n]
			n--
		} else {
			nums1[k] = nums1[m]
			nums1[m] = 0
			m--
		}
		k--
	}
	return nums1
}
