package main

import (
	"fmt"
	"sync"
)

type Job struct {
	ID     int
	Result int
}

func Worker(wg *sync.WaitGroup, in <-chan *Job, out chan<- *Job) {

	defer wg.Done()

	for val := range in {
		RunJob(val)
		out <- val
	}
}

func RunJob(job *Job) *Job {
	job.Result = job.ID * job.ID

	return job
}

func PrintResult(wg *sync.WaitGroup, out <-chan *Job) {

	defer wg.Done()

	for val := range out {
		fmt.Printf("job %d was processed. Result: %d \n", val.ID, val.Result)
	}
}

func main() {
	n := 100
	wg := &sync.WaitGroup{}
	outWg := &sync.WaitGroup{}

	out := make(chan *Job)

	in := make(chan *Job)

	outWg.Add(1)
	go PrintResult(outWg, out)

	wg.Add(5)

	for i := 0; i < 5; i++ {
		go Worker(wg, in, out)
	}

	for i := 0; i < n; i++ {
		in <- &Job{
			ID: i,
		}
	}

	close(in)
	wg.Wait()
	close(out)
	outWg.Wait()

}
