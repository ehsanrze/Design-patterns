package main

import (
	"context"
	"fmt"
	"time"

	"math/rand"

	"golang.org/x/sync/errgroup"
)

func Task(ctx context.Context, id int) error {
	var a int

	ticker := time.NewTicker(time.Second)
	defer ticker.Stop()

	for {
		select {
		case <-ticker.C:
			a = rand.Intn(100)

			if a < 10 {
				return fmt.Errorf("Job %d has failed \n", id)
			}

			fmt.Printf("Job %d has been processed \n", id)
		case <-ctx.Done():
			return fmt.Errorf("Job %d has failed \n", id)
		}
	}
}

func main() {

	g, ctx := errgroup.WithContext(context.Background())

	for i := 1; i <= 3; i++ {
		g.Go(func() error {
			return Task(ctx, i)
		})
	}


	if err := g.Wait(); err != nil {
		fmt.Println(err)
	}

}
