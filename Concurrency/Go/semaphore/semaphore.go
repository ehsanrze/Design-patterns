package main

import (
	"fmt"
	"sync"
	"time"
)

type Job struct {
	ID int
}

func NewSemaphore(n int) *Semaphore {
	s := &Semaphore{
		ch: make(chan struct{}, n),
	}

	for i := 0; i < n; i++ {
		s.Release()
	}

	return s
}

type Semaphore struct {
	ch chan struct{}
}

func (s *Semaphore) Wait() {
	<-s.ch
}

func (s *Semaphore) Release() {
	s.ch <- struct{}{}
}

func RunJob(job *Job) {
	fmt.Printf("Job %d starts... \n", job.ID)
	time.Sleep(1 * time.Second)
	fmt.Printf("Job %d finished... \n", job.ID)
}

func Worker(wg *sync.WaitGroup, s *Semaphore, job *Job) {

	defer wg.Done()

	s.Wait()
	defer s.Release()

	RunJob(job)
}

func main() {
	n := 100
	s := NewSemaphore(5)
	wg := &sync.WaitGroup{}

	wg.Add(n)

	for i := 0; i < n; i++ {
		go Worker(wg, s, &Job{
			ID: i,
		})
	}

	wg.Wait()
}
