package main

import "fmt"

func main() {
	ch := make(chan int)
	go generateEven(10, ch)
	for v := range ch {
		fmt.Println(v)
	}
}

func generateEven(n int, ch chan int) {
	for i := 2; i <= n; i += 2 {
		ch <- i
	}
	close(ch)
}
