package main

/*////////////////// Funcs as values ///////////////////

func compute(fn func(float64, float64) float64) float64 {
	return fn(3, 4)
}

func main2() {
	hypot := func(x, y float64) float64 {
		return math.Sqrt(x*x + y*y)
	}
	fmt.Println(hypot(5, 12))

	fmt.Println(compute(hypot))
	fmt.Println(compute(math.Pow))
}

// //////////////////// Closure /////////////////
func adder() func(int) int {
	sum := 0
	return func(x int) int {
		sum += x
		return sum
	}
}

func main() {
	// two different adders operating independently on their own sums
	pos, neg := adder(), adder()

	for i := 0; i < 10; i++ {
		fmt.Println(
			pos(i),
			neg(-2*i),
		)
	}
}

/////////////////////////////////////////////////*/
