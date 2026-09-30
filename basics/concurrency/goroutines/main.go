package main

import (
	"fmt"
	"sync"
)

func main() {
	count := 0
	var wg sync.WaitGroup

	for i := 0; i < 100; i++ {

		wg.Add(1)
		go func(wg *sync.WaitGroup) {
			var mu sync.Mutex
			wg.Done()
			mu.Lock()
			count++
			defer mu.Unlock()
		}(&wg)
	}

	wg.Wait()

	fmt.Println(count)
}
