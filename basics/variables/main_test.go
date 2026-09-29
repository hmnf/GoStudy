package main

import "testing"

func TestSumDigits(t *testing.T) {
	tests := []struct {
		num  int
		want int
	}{
		{1234, 10},
		{105, 6},
		{5, 5},
		{0, 0},
	}

	for _, test := range tests {
		got := sumDigits(test.num)

		if got != test.want {
			t.Errorf(
				"sumDigits(%d) = %d; want %d",
				test.num,
				got,
				test.want,
			)
		}
	}
}

func TestMaxOfThree(t *testing.T) {
	tests := []struct {
		nums []float64
		want float64
	}{
		{[]float64{5, 12, 8}, 12},
		{[]float64{100, 5, 3}, 100},
		{[]float64{1, 2, 100}, 100},
		{[]float64{-5, -2, -10}, -2},
	}

	for _, test := range tests {
		got := maxOfThree(test.nums[0], test.nums[1], test.nums[2])

		if got != test.want {
			t.Errorf(
				"maxOfThree(%v) = %.2f, want %.2f",
				test.nums,
				got,
				test.want,
			)
		}
	}
}

func TestIsPalindrome(t *testing.T) {
	tests := []struct {
		num  int
		want bool
	}{
		{1221, true},
		{12321, true},
		{1234, false},
		{-121, false},
		{0, true},
		{7, true},
	}

	for _, test := range tests {
		got := isPalindrome(test.num)

		if got != test.want {
			t.Errorf(
				"isPalindrome(%d) = %v, want %v",
				test.num,
				got,
				test.want,
			)
		}
	}
}

func TestCalc(t *testing.T) {
	tests := []struct {
		name     string
		num1     float64
		num2     float64
		operator string
		want     float64
		wantErr  bool
	}{
		{"addition", 10, 5, "+", 15, false},
		{"subtraction", 10, 5, "-", 5, false},
		{"multiplication", 10, 5, "*", 50, false},
		{"division", 10, 4, "/", 2.5, false},
		{"division by zero", 10, 0, "/", 0, true},
		{"invalid operator", 10, 5, "%", 0, true},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			got, err := calc(
				test.num1,
				test.num2,
				test.operator,
			)

			if test.wantErr {
				if err == nil {
					t.Errorf("expected error, got nil")
				}
				return

			}

			if err != nil {
				t.Errorf("unexpected error: %v", err)
				return
			}

			if got != test.want {
				t.Errorf(
					"calc(%.2f %s %.2f) = %.2f; want %.2f",
					test.num1,
					test.operator,
					test.num2,
					got,
					test.want,
				)
			}
		})
	}
}
