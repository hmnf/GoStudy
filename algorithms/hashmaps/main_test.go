package main

import "testing"

func TestIsAnagram(t *testing.T) {
	tests := []struct {
		name string
		s    string
		t    string
		want bool
	}{
		{
			name: "anagram",
			s:    "anagram",
			t:    "nagaram",
			want: true,
		}, {
			name: "no anagram",
			s:    "cat",
			t:    "rat",
			want: false,
		}, {
			name: "russian",
			s:    "кот",
			t:    "ток",
			want: true,
		}, {
			name: "-1 in count",
			s:    "ab",
			t:    "abc",
			want: false,
		}, {
			name: "nothing",
			s:    "",
			t:    "",
			want: true,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			got := isAnagram(test.s, test.t)

			if got != test.want {
				t.Errorf(
					"isAnagram(%v,%v) = %v; want %v",
					test.s,
					test.t,
					got,
					test.want,
				)
			}
		})
	}
}
