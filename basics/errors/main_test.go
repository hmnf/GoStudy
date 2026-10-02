package main

import (
	"errors"
	"testing"
)

func TestDivide(t *testing.T) {
	tests := []struct {
		name    string
		a       float64
		b       float64
		want    float64
		wantErr error
	}{
		{
			"1",
			10,
			2,
			5,
			nil,
		},
		{
			"2",
			-10,
			2,
			-5,
			nil,
		},
		{
			"3",
			10,
			0,
			0,
			errors.New("divizion by zero"),
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			got, err := Divide(test.a, test.b)
			if !errors.Is(err, test.wantErr) {
				t.Fatalf(
					"Divide returned %v, wanted %v",
					err,
					test.wantErr,
				)
			}

			if got != test.want {
				t.Errorf(
					"Divide(%.2f,%.2f) = %.2f, want %.2f",
					test.a,
					test.b,
					got,
					test.want,
				)
			}
		})
	}
}
