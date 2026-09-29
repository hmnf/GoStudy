package main

func TwoSum(nums []int, target int) []int {
	set := make(map[int]int)

	for v, num := range nums {
		if _, exists := set[num]; !exists {
			set[num] = v
			if _, exists = set[target-num]; exists {
				return []int{set[target-num], v}
			}
		}
	}

	return []int{}
}
