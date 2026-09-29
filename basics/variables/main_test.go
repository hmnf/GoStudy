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
		num1     float64
		num2     float64
		operator string
		want     float64
	}{
		{10, 5, "+", 15},
		{10, 5, "-", 5},
		{10, 5, "*", 50},
		{10, 4, "/", 2.5},
	}

	for _, test := range tests {
		got, err := calc(test.num1, test.num2, test.operator)
		if err != nil {
			t.Errorf("unexpected error: %v", err)
		}

		if got != test.want {
			t.Errorf(
				"calc(%.2f %v %.2f) = %.2f, want %.2f",
				test.num1,
				test.operator,
				test.num2,
				got,
				test.want,
			)
		}
	}

	_, err := calc(10, 0, "/")

	if err == nil {
		t.Errorf("divizion by zero, but nil error")
	}

	_, err = calc(10, 5, "%")
	if err == nil {
		t.Errorf("wrong operator, but nil error")
	}
}
