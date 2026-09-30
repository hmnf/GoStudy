package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	go worker()
}

func worker() {
	ctx, cancel := context.WithTimeout(
		context.Background(),
		1*time.Second,
	)
	defer cancel()
	for {
		time.Sleep(300 * time.Millisecond)
		select {
		case <-ctx.Done():
			fmt.Println("working canceled")
		default:
			fmt.Println("Working...")
		}
	}
}
