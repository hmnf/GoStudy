package main

func TwoSum(nums []int, target int) []int {
	seen := make(map[int]int)

	for i, num := range nums {
		if _, exists := seen[target-num]; exists {
			return []int{seen[target-num], i}
		}
		seen[num] = i
	}

	return []int{}
}
