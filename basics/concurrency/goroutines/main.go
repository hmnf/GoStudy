package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	msg := make(chan string)

	go worker(msg)
}

func worker(msg chan string) {
	ctx, cancel := context.WithTimeout(
		context.Background(),
		1*time.Second,
	)
	defer cancel()
	for {
		msg <- "Working..."
		time.Sleep(300 * time.Millisecond)
		select {
		case v := <-msg:
			fmt.Println(v)
		case <-ctx.Done():
			fmt.Println("working canceled")
		}
	}
}
