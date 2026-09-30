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

	go worker(msg, ctx)

	defer cancel()
}

func worker(msg chan string, ctx context.Context) {
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
