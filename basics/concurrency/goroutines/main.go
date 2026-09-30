package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	ctx, cancel := context.WithTimeout(
		context.Background(),
		1*time.Second,
	)

	msg := make(chan string)

	go worker(msg)
	select {
	case v := <-msg:
		fmt.Println(v)
	case <-ctx.Done():
		fmt.Println("working canceled")
	}

	defer cancel()
}

func worker(msg chan string) {
	i := 0
	for i < 1 {
		msg <- "Working..."
		time.Sleep(300 * time.Millisecond)
	}
}
