package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

func main() {
	ctx, cancel := context.WithTimeout(
		context.Background(),
		1*time.Second,
	)
	defer cancel()

	var wg sync.WaitGroup
	for {
		wg.Add(1)
		go worker(ctx, &wg)
	}
}

func worker(ctx context.Context, wg *sync.WaitGroup) {
	defer wg.Done()
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
