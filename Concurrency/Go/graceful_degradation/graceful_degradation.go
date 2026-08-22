package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

type Job struct {
	ID     int
	Result int
}

func Worker(ctx context.Context, wg *sync.WaitGroup, in <-chan *Job, out chan<- *Job) {

	defer wg.Done()

	for {
		select {
		case <-ctx.Done():
			return
		case val, ok := <-in:
			if !ok {
				return
			}

			RunJob(val)

			select {
			case <-ctx.Done():
				return
			case out <- val:
			}

		}
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

func Generate(ctx context.Context) <-chan *Job {

	out := make(chan *Job, 100)

	go func() {

		defer close(out)

		i := 0

		ticker := time.NewTicker(500 * time.Millisecond)
		defer ticker.Stop()

		for {
			select {
			case <-ctx.Done():
				return
			case <-ticker.C:

				job := &Job{
					ID: i,
				}

				select {
				case <-ctx.Done():
					return
				case out <- job:
					i++
				}
			}
		}
	}()

	return out
}

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	wg := &sync.WaitGroup{}
	outWg := &sync.WaitGroup{}

	out := make(chan *Job)
	in := Generate(ctx)

	outWg.Add(1)
	go PrintResult(outWg, out)

	wg.Add(5)

	for i := 0; i < 5; i++ {
		go Worker(ctx, wg, in, out)
	}

	wg.Wait()
	close(out)
	outWg.Wait()

}
