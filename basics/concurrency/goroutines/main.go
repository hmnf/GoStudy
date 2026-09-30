package main

import (
	"fmt"
	"time"
)

func main() {
	fast := make(chan string)
	slow := make(chan string)

	go msg1(fast)
	go msg2(slow)

	select {
	case v := <-fast:
		fmt.Println(v)

	case v := <-slow:
		fmt.Println(v)

	}
}

func msg1(fast chan string) {
	fast <- "fast finished"
}

func msg2(slow chan string) {
	time.Sleep(500 * time.Millisecond)
	slow <- "slow finished"
}
