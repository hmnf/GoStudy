package main

import (
	"errors"
	"fmt"
)

var ErrDivizionByZero error

func main() {
	ErrDivizionByZero = errors.New("divizion by zero")
	res, err := Divide(10, 2)

	if errors.Is(err, ErrDivizionByZero) {
		fmt.Println(err)
		return
	}

	fmt.Println(res)
}

func Divide(a, b float64) (float64, error) {
	if b == 0 {
		return 0, errors.New("divizion by zero")
	}

	return a / b, nil
}
