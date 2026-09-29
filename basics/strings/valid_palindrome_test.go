package main

import "testing"

func TestIsPalindrome(t *testing.T) {
	tests := []struct {
		name string
		s    string
		want bool
	}{
		{
			name: "palindrome",
			s:    "racecar",
			want: true,
		},
		{
			name: "not palindrome",
			s:    "hello",
			want: false,
		},
		{
			name: "1 letter",
			s:    "a",
			want: true,
		},
		{
			name: "empty",
			s:    "",
			want: true,
		},
		{
			name: "unicode palindrome",
			s:    "топот",
			want: true,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			got := IsPalindrome(test.s)

			if got != test.want {
				t.Errorf(
					"IsPalindrome(%v) = %v; want %v",
					test.s,
					got,
					test.want,
				)
			}
		})
	}
}
