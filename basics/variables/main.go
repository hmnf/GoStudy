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
			var num1, num2 float64
			var operator string

			fmt.Print("Enter two numbers: ")
			if _, err := fmt.Scan(&num1, &num2); err != nil {
				fmt.Println("Error:", err)
				return
			}

			fmt.Print("Enter operator (+,-,*,/): ")
			if _, err := fmt.Scan(&operator); err != nil {
				fmt.Println("Error:", err)
				return
			}

			result, err := calc(num1, num2, operator)
			if err != nil {
				fmt.Println("Error:", err)
				return
			}
			fmt.Printf("Result: %.2f\n", result)
		},
		"sumDigits": func() {
			var num int
			fmt.Print("Enter a number: ")
			if _, err := fmt.Scan(&num); err != nil {
				fmt.Println("Error:", err)
				return
			}
			fmt.Println("Result:", sumDigits(num))
		},
		"isPalindrome": func() {
			var num int
			fmt.Print("Enter a number: ")
			if _, err := fmt.Scan(&num); err != nil {
				fmt.Println("Error:", err)
				return
			}
			fmt.Println("Result:", isPalindrome(num))
		},
		"maxOfThree": func() {
			var num1, num2, num3 float64
			fmt.Print("Enter three numbers: ")
			if _, err := fmt.Scan(&num1, &num2, &num3); err != nil {
				fmt.Println("Error:", err)
				return
			}
			fmt.Println("Result:", maxOfThree(num1, num2, num3))
		},
	}

	f, ok := functions[funcName]
	if !ok {
		fmt.Println("Function not found")
		return
	}

	f()
}

func calc(num1, num2 float64, operator string) (float64, error) {
	switch operator {
	case "+":
		return num1 + num2, nil
	case "-":
		return num1 - num2, nil
	case "*":
		return num1 * num2, nil
	case "/":
		if num2 == 0 {
			return 0, fmt.Errorf("division by zero")
		}

		return num1 / num2, nil
	default:
		return 0, fmt.Errorf("invalid operator")
	}
}

func sumDigits(num int) int {
	var sum int

	for num > 0 {
		sum += num % 10
		num /= 10
	}

	return sum
}

func isPalindrome(num int) bool {
	if num < 0 {
		return false
	}

	l := numLen(num)

	for i := 0; i < l/2; i++ {
		if num/int(math.Pow(10, float64(i)))%10 != num/int(math.Pow(10, float64(l-i-1)))%10 {
			return false
		}
	}

	return true
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

func maxOfThree(num1, num2, num3 float64) float64 {
	m := num1

	if num2 > m {
		m = num2
	}
	if num3 > m {
		m = num3
	}

	return m
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
