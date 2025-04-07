package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())

	number := make([]int, 101)
	for i := range number {
		number[i] = i
	}

	respCh := make(chan int)
	resultCh := make(chan int, 10)
	var wg sync.WaitGroup
	wg.Add(2)

	go func() {
		defer wg.Done()
		defer close(respCh)

		rand.Shuffle(len(number), func(i, j int) {
			number[i], number[j] = number[j], number[i]
		})
		for _, num := range number[:10] {
			respCh <- num
		}
	}()

	go func() {
		defer wg.Done()
		defer close(resultCh)
		for i := 0; i < 10; i++ {
			num := <-respCh
			resultCh <- num * num
		}
	}()

	for res := range resultCh {
		fmt.Println(res)
	}
	wg.Wait()
}
