package main

import "testing"

func TestIsValid(t *testing.T) {
	tests := []struct {
		name string
		s    string
		want bool
	}{
		{
			name: "1",
			s:    "()",
			want: true,
		},
		{
			name: "2",
			s:    "()[]{}",
			want: true,
		},
		{
			name: "3",
			s:    "([)]",
			want: false,
		},
		{
			name: "4",
			s:    "{[]}",
			want: true,
		},
		{
			name: "5",
			s:    "([{(){}}])({[()]})",
			want: true,
		},
		{
			name: "6",
			s:    "(([{(){}}])({[()])}))",
			want: false,
		},
		{
			name: "7",
			s:    "",
			want: true,
		},
		{
			name: "8",
			s:    "(]",
			want: false,
		},
		{
			name: "9",
			s:    ")",
			want: false,
		},
		{
			name: "10",
			s:    "(((",
			want: false,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			got := isValid(test.s)

			if got != test.want {
				t.Fatalf(
					"isValid(%v) = %v, want %v",
					test.s,
					got,
					test.want,
				)
			}
		})
	}
}
