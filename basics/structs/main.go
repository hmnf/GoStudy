package main

import "fmt"

type BankAccount struct {
	Owner   string
	balance float64
}

func main() {
	b := BankAccount{
		Owner: "Ars",
	}

	b.Deposit(1000)
	b.Withdraw(700)
	fmt.Println(b.Balance())
}

func (b *BankAccount) Deposit(amount float64) {
	if amount > 0 {
		b.balance += amount
	}
}

func (b *BankAccount) Withdraw(amount float64) bool {
	if amount <= 0 || amount > b.balance {
		return false
	}
	b.balance -= amount
	return true
}

func (b BankAccount) Balance() float64 {
	return b.balance
}
