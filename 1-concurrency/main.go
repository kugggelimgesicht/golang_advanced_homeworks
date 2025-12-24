package main

import (
	"fmt"
	"math/rand"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	numbersCh := make(chan int, 10)
	squaresCh := make(chan int, 10)
	wg.Add(2)
	go func() {
		defer wg.Done()
		defer close(numbersCh)

		randomNumbers := make([]int, 10)
		for i := range randomNumbers {
			randomNumbers[i] = rand.Intn(100)
		}
		for i := range randomNumbers {
			numbersCh <- randomNumbers[i]
		}
	}()
	go func() {
		defer wg.Done()
		defer close(squaresCh)
		for n := range numbersCh {
			squaresCh <- n * n
		}

	}()

	wg.Wait()

	for s := range squaresCh {
		fmt.Print(s, ' ')
	}
}
