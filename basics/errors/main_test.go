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
			ErrDivizionByZero,
		},
		{
			"4",
			0,
			5,
			0,
			nil,
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

func TestGetUserName(t *testing.T) {
	tests := []struct {
		name    string
		id      int
		want    string
		wantErr error
	}{
		{
			name:    "1",
			id:      1,
			want:    "Arseniy",
			wantErr: nil,
		},
		{
			name:    "2",
			id:      2,
			want:    "Ivan",
			wantErr: nil,
		},
		{
			name:    "3",
			id:      10,
			want:    "",
			wantErr: ErrUserNotFound,
		},
		{
			name:    "4",
			id:      -1,
			want:    "",
			wantErr: ErrUserNotFound,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			got, err := GetUserName(test.id)

			if !errors.Is(err, test.wantErr) {
				t.Fatalf(
					"wanted error %v != got error %v",
					test.wantErr,
					err,
				)
			}

			if got != test.want {
				t.Errorf(
					"GetUserName(%d) = %v, want %v",
					test.id,
					got,
					test.want,
				)
			}
		})
	}
}

func TestValidateUser(t *testing.T) {
	tests := []struct {
		testname string
		name     string
		email    string
		want     error
	}{
		{
			testname: "no empty fields",
			name:     "Ars",
			email:    "kkkk",
			want:     nil,
		},
		{
			testname: "empty name",
			name:     "",
			email:    "kkkk",
			want: ValidationError{
				Field: "name",
			},
		},
		{
			testname: "empty email",
			name:     "Ars",
			email:    "",
			want: ValidationError{
				Field: "email",
			},
		},
		{
			testname: "all empty ",
			name:     "",
			email:    "",
			want: ValidationError{
				Field: "name",
			},
		},
	}

	var validationErr ValidationError

	for _, test := range tests {
		t.Run(test.testname, func(t *testing.T) {
			got := ValidateUser(test.name, test.email)

			if errors.As(got, &validationErr) {
				if validationErr.Field != test.want.Field {
					t.Errorf(
						"wanted error %v != got error %v",
						test.want,
						got,
					)
				}
			}
		})
	}
}
