package main

import (
	"fmt"
	"math/rand"
	"sync"
)

func WorkerPool[T any](n int, in <-chan T, action func(a T) T) <-chan T {

	wg := sync.WaitGroup{}
	wg.Add(n)

	out := make(chan T)

	go func() {
		for i := 0; i < n; i++ {
			go Worker[T](&wg, in, out, action)
		}

		wg.Wait()
		close(out)
	}()

	return out
}

func Worker[T any](wg *sync.WaitGroup, in <-chan T, out chan<- T, action func(a T) T) {

	defer wg.Done()

	for val := range in {
		out <- action(val)
	}
}

func Generate(n int) <-chan int {

	out := make(chan int)

	go func() {
		defer close(out)

		for i := 0; i < n; i++ {
			out <- rand.Intn(100)
		}
	}()

	return out
}

func FinalResult(in <-chan int) {
	for val := range in {
		fmt.Println(val)
	}
}

func main() {
	ch1 := Generate(100)

	ch2 := WorkerPool[int](5, ch1, func(a int) int {
		return a * a
	})

	FinalResult(ch2)
}
