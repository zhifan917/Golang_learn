package main

import (
	"context"
	"fmt"
	"time"
)

func worker(ctx context.Context) {
LOOP:
	for {
		fmt.Printf("worker\n")
		time.Sleep(time.Second)

		select {
		case <-ctx.Done():
			break LOOP
		default:
		}
	}
}

func main() {
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)

	go worker(ctx)
	time.Sleep(time.Second * 3)
	cancel()
}
