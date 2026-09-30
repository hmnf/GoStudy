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

	go worker(ctx)

	defer cancel()
}

func worker(ctx context.Context) {
	defer ctx.Done()
	for i := 0; i < 10000000000000; i++ {
		fmt.Println("Working...")
		time.Sleep(300 * time.Millisecond)
	}
}
