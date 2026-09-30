package main

import "testing"

func TestValidPalindrome(t *testing.T) {
	tests := []struct {
		name string
		s    string
		want bool
	}{
		{
			name: "basic",
			s:    "A man, a plan, a canal: Panama",
			want: true,
		},
		{
			name: "false",
			s:    "race a car",
			want: false,
		},
		{
			name: "space",
			s:    " ",
			want: true,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			got := ValidPalindrome(test.s)

			if got != test.want {
				t.Errorf(
					"ValidPalindrome(%s) = %v, want %v",
					test.s,
					got,
					test.want,
				)
			}
		})
	}
}
