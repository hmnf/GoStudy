package main

import (
	"errors"
	"fmt"
)

var ErrDivizionByZero = errors.New("divizion by zero")

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
