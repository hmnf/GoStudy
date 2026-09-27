package main

import "fmt"

// //////////////// Training ////////////////////////
var num1, num2 float64

func main() {
	fmt.Print(num1, num2)

	_, err := fmt.Scan(&num1, &num2)
	if err != nil {
		fmt.Println("invalid input")
		return
	}

	fmt.Printf("Sum: %f\n", num1+num2)
	fmt.Printf("Diff: %f\n", num1-num2)
	fmt.Printf("Product: %f\n", num1*num2)
	fmt.Printf("Division: %2.2f\n", num1/num2)
}

// //////////// Structs //////////////
type Vertex struct {
	X int
	Y int
}

func main2() {
	fmt.Println(Vertex{1, 2})
}

/////////////////////////////////////
