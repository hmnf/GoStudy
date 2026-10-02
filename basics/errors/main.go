package main

import (
	"errors"
	"fmt"
)

var (
	ErrDivizionByZero = errors.New("Divizion by zero")
	ErrUserNotFound   = errors.New("user not found")
)

func main() {
	res, err := Divide(10, 2)

	if errors.Is(err, ErrDivizionByZero) {
		fmt.Println(err)
		return
	}

	fmt.Println(res)
}

func Divide(a, b float64) (float64, error) {
	if b == 0 {
		return 0, ErrDivizionByZero
	}

	return a / b, nil
}

func GetUserName(id int) (string, error) {
	switch id {
	case 1:
		return "Arseniy", nil
	case 2:
		return "Ivan", nil
	default:
		return "", fmt.Errorf("User with id %d: %w", id, ErrUserNotFound)
	}
}
