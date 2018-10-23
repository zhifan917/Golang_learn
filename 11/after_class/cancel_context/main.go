package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

var wg sync.WaitGroup

func worker2(ctx context.Context) {
LOOP:
	for {
		fmt.Printf("worker2\n")
		time.Sleep(time.Second)
		select {
		case <-ctx.Done():
			break LOOP
		default:

		}
	}
}

func worker(ctx context.Context) {
	//ctx2, cancel := context.WithCancel(ctx)
	go worker2(ctx)
	//cancel()

LOOP:
	for {
		fmt.Printf("worker\n")
		time.Sleep(time.Second)
		select {
		case <-ctx.Done(): //Done函数本身是一个channel
			break LOOP
		default:

		}
	}

	wg.Done()
}

func main() {
	ctx := context.Background()            //生成一个context，类似于根context
	ctx, cancel := context.WithCancel(ctx) //会返回一个子context和一个取消函数
	wg.Add(1)
	go worker(ctx)
	time.Sleep(time.Second * 3)
	cancel() //要取消该goroutine时，直接调用cancel函数(往channel中放一个关键字东西进去，然后Done函数（本质是channel）读取到该关键字，该goroutine就优雅退出了，go官方关于该原理的介绍https://blog.golang.org/context)
	wg.Wait()
}
