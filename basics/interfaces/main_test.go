package main

import "testing"

func TestSpeaker(t *testing.T) {
	p := Person{
		Name: "hmnf",
	}

	got := p.Speak()

	want := "My name is hmnf"

	if got != want {
		t.Fatalf(
			"Speak() = %v, want %v",
			got,
			want,
		)
	}

	d := Dog{}

	got = d.Speak()
	want = "woof woof"

	if got != want {
		t.Fatalf(
			"Speak() = %v, want %v",
			got,
			want,
		)
	}
}
