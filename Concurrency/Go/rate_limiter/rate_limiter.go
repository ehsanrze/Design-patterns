package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

type RateLimiter struct {
	count int
	ch    chan struct{}
	ctx   context.Context
}

func NewRateLimiter(ctx context.Context, count int, timeout time.Duration) *RateLimiter {

	r := &RateLimiter{
		count: count,
		ch:    make(chan struct{}, count),
		ctx:   ctx,
	}

	r.Reset()
	go r.Timer(timeout)

	return r
}

func (r *RateLimiter) Timer(timeout time.Duration) {
	ticker := time.NewTicker(timeout)

	defer ticker.Stop()

	for {
		select {
		case <-ticker.C:
			r.Reset()
		case <-r.ctx.Done():
			return
		}
	}
}

func (r *RateLimiter) Reset() {
drain:
	for {
		select {
		case <-r.ch:
			// remove value
		default:
			break drain
		}
	}

	for i := 0; i < r.count; i++ {
		r.ch <- struct{}{}
	}
}

func (r *RateLimiter) Wait() bool {
	select {
	case <-r.ch:
		return true
	case <-r.ctx.Done():
		return false
	}
}

type Job struct {
	ID int
}

func CallAPI(job *Job, r *RateLimiter, w *sync.WaitGroup) {
	defer w.Done()

	if !r.Wait() {
		return
	}

	fmt.Printf("Calling API for job %d\n", job.ID)
}

func main() {
	wg := sync.WaitGroup{}
	wg.Add(20)

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	r := NewRateLimiter(ctx, 5, time.Second)

	for i := 0; i < 20; i++ {
		go CallAPI(&Job{
			ID: i,
		}, r, &wg)
	}

	wg.Wait()
}
