package main

import "testing"

func testSum(t *testing.T) {
	tests := []struct {
		nums []int
		want int
	}{
		{[]int{1, 2, 3, 4}, 10},
		{[]int{-1, -2, -3}, -6},
		{[]int{}, 0},
	}
	for _, test := range tests {
		got := sum(test.nums)
		if got != test.want {
			t.Errorf(
				"sum(%v) = %d; want %d",
				test.nums,
				got,
				test.want,
			)
		}
	}
}

func testMax(t *testing.T) {
	tests := []struct {
		nums []int
		want int
	}{
		{[]int{1, 2, 3, 4}, 4},
		{[]int{-1, -2, -3}, -1},
		{[]int{5}, 5},
	}
	for _, test := range tests {
		got := max(test.nums)
		if got != test.want {
			t.Errorf(
				"max(%v) = %d; want %d",
				test.nums,
				got,
				test.want,
			)
		}
	}
}

func testContains(t *testing.T) {
	tests := []struct {
		nums   []int
		target int
		want   bool
	}{
		{[]int{4, 7, 10}, 7, true},
		{[]int{4, 7, 10}, 20, false},
	}
	for _, test := range tests {
		got := contains(test.nums, test.target)
		if got != test.want {
			t.Errorf(
				"contains(%v, %d) = %t; want %t",
				test.nums,
				test.target,
				got,
				test.want,
			)
		}
	}
}

func testOnlyEven(t *testing.T) {
	tests := []struct {
		nums []int
		want []int
	}{
		{[]int{1, 2, 3, 4, 5, 6}, []int{2, 4, 6}},
		{[]int{1, 3, 5}, []int{}},
		{[]int{2, 4, 6}, []int{2, 4, 6}},
	}
	for _, test := range tests {
		got := onlyEven(test.nums)
		if !equalSlices(got, test.want) {
			t.Errorf(
				"onlyEven(%v) = %v; want %v",
				test.nums,
				got,
				test.want,
			)
		}

	}
}

func testReverse(t *testing.T) {
	tests := []struct {
		nums []int
		want []int
	}{
		{[]int{1, 2, 3, 4, 5}, []int{5, 4, 3, 2, 1}},
		{[]int{1}, []int{1}},
		{[]int{}, []int{}},
	}
	for _, test := range tests {
		got := reverse(test.nums)
		if !equalSlices(got, test.want) {
			t.Errorf(
				"reverse(%v) = %v; want %v",
				test.nums,
				got,
				test.want,
			)
		}
	}
}

func equalSlices(a, b []int) bool {
	if len(a) != len(b) {
		return false
	}
	for i := 0; i < len(a); i++ {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}
