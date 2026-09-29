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
	if b.balance-amount < 0 || amount < 0 {
		return false
	}
	b.balance -= amount
	return true
}

func (b BankAccount) Balance() float64 {
	return b.balance
}

/*
type User struct {
	Name string
	Age  int
}

func main() {
	user := User{
		Name: "Ars",
		Age:  17,
	}
	user.Birthday()
	fmt.Println(user.IsAdult())
}

func (u User) GetName() string {
	return u.Name
}

func (u User) IsAdult() bool {
	return u.Age >= 18
}

func (u *User) Birthday() {
	u.Age++
}
*/
