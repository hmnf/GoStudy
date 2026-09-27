package main

///////// SQUARE ROOT FUNCTION ////////////

import (
	"fmt"
	"math"
)

func Sqrt(x float64) float64 {
	var z float64
	z = 1
	for math.Abs(z*z-x) > 0.01 {
		// Newton's method for aproxing square root
		z -= (z*z - x) / (2 * z)
		fmt.Println(z)
	}
	return z
}

func main() {
	fmt.Println(Sqrt(4))
}

//////////// Range  //////////////////

var pow = []int{1, 2, 4, 8, 16, 32, 64, 128}

func main2() {
	for i, v := range pow {
		fmt.Printf("2**%d = %d\n", i, v)
	}
}

//////////////////////////////////////
