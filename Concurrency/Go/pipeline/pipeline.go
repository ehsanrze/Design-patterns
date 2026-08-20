package main

import (
	"fmt"
	"math/rand"
)

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

func Filter(in <-chan int) <-chan int {

	out := make(chan int)

	go func() {
		defer close(out)

		for val := range in {
			if val%2 == 0 {
				out <- val
			}
		}
	}()

	return out
}

func Square(in <-chan int) <-chan int {

	out := make(chan int)

	go func() {
		defer close(out)

		for val := range in {
			out <- val * val
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
	ch2 := Filter(ch1)
	ch3 := Square(ch2)
	FinalResult(ch3)
}
