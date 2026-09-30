package main

import "fmt"

func square(n int, ch chan int) {
	ch <- n * n
}

func main() {
	ch := make(chan int)
	go square(2, ch)
	go square(3, ch)
	go square(4, ch)

	for i := 0; i < 3; i++ {

		v := <-ch
		fmt.Println(v)
	}
}
