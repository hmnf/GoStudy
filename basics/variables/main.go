package main

import (
	"fmt"
	"math"
)

// //////////////// Training ////////////////////////

func main() {
	var funcName string

	fmt.Print("Enter function name(calc, sumDigits, isPalindrome, maxOfThree): ")

	_, err := fmt.Scan(&funcName)
	if err != nil {
		fmt.Println("invalid input")
		return
	}

	functions := map[string]func(){
		"calc": func() {
			result, err := calc()
			if err != nil {
				fmt.Println("Error:", err)
				return
			}
			fmt.Println("Result:", result)
		},
		"sumDigits": func() {
			result, err := sumDigits()
			if err != nil {
				fmt.Println("Error:", err)
				return
			}
			fmt.Println("Result:", result)
		},
		"isPalindrome": func() {
			result, err := isPalindrome()
			if err != nil {
				fmt.Println("Error:", err)
				return
			}
			fmt.Println("Result:", result)
		},
		"maxOfThree": func() {
			result, err := maxOfThree()
			if err != nil {
				fmt.Println("Error:", err)
				return
			}
			fmt.Println("Result:", result)
		},
	}

	f, ok := functions[funcName]
	if !ok {
		fmt.Println("Function not found")
		return
	}

	f()
}

func calc() (float64, error) {
	var num1, num2 float64
	var operator string

	fmt.Print("Enter two numbers:")
	_, err := fmt.Scan(&num1, &num2)
	if err != nil {
		return 0, err
	}

	fmt.Print("Enter operator (+,-,*,/): ")
	_, err = fmt.Scan(&operator)
	if err != nil {
		return 0, err
	}

	switch operator {
	case "+":
		return num1 + num2, nil
	case "-":
		return num1 - num2, nil
	case "*":
		return num1 * num2, nil
	case "/":
		if num2 == 0 {
			fmt.Println("Division by zero is not allowed")
			return 0, fmt.Errorf("division by zero")
		}

		return num1 / num2, nil
	default:
		return 0, fmt.Errorf("invalid operator")
	}
}

func sumDigits() (int, error) {
	var num int
	fmt.Print("Enter a number: ")
	_, err := fmt.Scan(&num)
	if err != nil {
		return 0, err
	}

	var sum int

	for num > 0 {
		sum += num % 10
		num /= 10
	}

	return sum, nil
}

func isPalindrome() (bool, error) {
	var num int

	fmt.Print("Enter a number: ")
	_, err := fmt.Scan(&num)
	if err != nil {
		return false, err
	}

	if num < 0 {
		return false, nil
	}

	l := numLen(num)

	for i := 0; i < l/2; i++ {
		if num/int(math.Pow(10, float64(i)))%10 != num/int(math.Pow(10, float64(l-i-1)))%10 {
			return false, nil
		}
	}

	return true, nil
}

func numLen(num int) int {
	if num == 0 {
		return 1
	}

	l := 0

	for num > 0 {
		l++
		num /= 10
	}
	return l
}

func maxOfThree() (float64, error) {
	var num1, num2, num3 float64

	fmt.Print("Enter three numbers: ")

	_, err := fmt.Scan(&num1, &num2, &num3)
	if err != nil {
		return 0, err
	}

	m := num1

	if num2 > m {
		m = num2
	}
	if num3 > m {
		m = num3
	}

	return m, nil
}

/* //////////// Structs //////////////
type Vertex struct {
	X int
	Y int
}

func main2() {
	fmt.Println(Vertex{1, 2})
}

*/
