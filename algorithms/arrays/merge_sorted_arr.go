package main

func MergeSortedArrays(nums1 []int, m int, nums2 []int, n int) []int {
	k := len(nums1) - 1
	n--
	m--
	for k >= 0 && m >= 0 && n >= 0 {
		if nums2[n] > nums1[m] {
			nums1[k] = nums2[n]
			n--
		} else {
			nums1[k] = nums1[m]
			m--
		}
		k--
	}
	for n >= 0 {
		nums1[k] = nums2[n]
		k--
		n--
	}
	return nums1
}
