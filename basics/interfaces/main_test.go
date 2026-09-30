package main

import "testing"

func TestSpeaker(t *testing.T) {
	p := Person{
		Name: "hmnf",
	}

	got := GetSpeech(p)

	want := "My name is hmnf"

	if got != want {
		t.Fatalf(
			"Speak() = %v, want %v",
			got,
			want,
		)
	}

	d := Dog{}

	got = GetSpeech(d)
	want = "woof woof"

	if got != want {
		t.Fatalf(
			"Speak() = %v, want %v",
			got,
			want,
		)
	}
}
