package main

import "testing"

func TestDeposit(t *testing.T) {
	account := BankAccount{}

	account.Deposit(1000)

	got := account.Balance()
	want := 1000.0

	if got != want {
		t.Errorf(
			"got %v, want %v",
			got,
			want,
		)
	}
}

func TestWithdraw(t *testing.T) {
	account := BankAccount{}

	account.Deposit(1000)

	ok := account.Withdraw(300)

	if !ok {
		t.Fatalf(
			"Withdraw returned false, want true",
		)
	}

	got := account.Balance()
	want := 700.0

	if got != want {
		t.Errorf(
			"got %v, want %v",
			got,
			want,
		)
	}
}

func TestWithdrawNotEnoughMoney(t *testing.T) {
	account := BankAccount{}

	account.Deposit(500)

	ok := account.Withdraw(1000)

	if ok {
		t.Fatalf(
			"Withdraw returned true, want false",
		)
	}

	got := account.Balance()
	want := 500.0

	if got != want {
		t.Errorf(
			"got %v, want %v",
			got,
			want,
		)
	}
}

func TestBalance(t *testing.T) {
	account := BankAccount{}

	got := account.Balance()
	want := 0.0

	if got != want {
		t.Errorf(
			"got %v, want %v",
			got,
			want,
		)
	}
}

func TestNegativeDeposit(t *testing.T) {
	account := BankAccount{}

	account.Deposit(-500)

	got := account.Balance()
	want := 0.0

	if got != want {
		t.Errorf(
			"got %v, want %v",
			got,
			want,
		)
	}
}

func TestNegativeWithdraw(t *testing.T) {
	account := BankAccount{}

	account.Deposit(500)

	ok := account.Withdraw(-200)

	if ok {
		t.Fatalf("expected false, returned true")
	}

	got := account.Balance()
	want := 500.0

	if got != want {
		t.Errorf(
			"got %v, want %v",
			got,
			want,
		)
	}
}
