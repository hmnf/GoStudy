package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(containsNearbyDuplicate([]int{1, 2, 3, 1}, 3))
}

func containsNearbyDuplicate(nums []int, k int) bool {
	rs := make(map[int][]int)

	for key, num := range nums {
		rs[num] = append(rs[num], key)
	}

	for _, arr := range rs {
		if len(arr) >= 2 {
			for i := 0; i < len(arr)-1; i++ {
				if math.Abs(float64(arr[i]-arr[i+1])) <= float64(k) {
					return true
				}
			}
		}
	}

	return false
}
