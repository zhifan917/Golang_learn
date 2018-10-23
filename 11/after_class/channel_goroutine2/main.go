package main

import (
	"fmt"
	"sync"
	"time"
)

var wg sync.WaitGroup

var exit bool

func worker(exitChan chan struct{}) { //struct{}表示空结构体
LOOP:
	for {
		fmt.Printf("worker\n")
		time.Sleep(time.Second)
		select {
		case <-exitChan:
			break LOOP
		default:

		}

	}

	wg.Done()
}

func main() {

	var exitChan chan struct{} = make(chan struct{})
	wg.Add(1)
	go worker(exitChan)
	time.Sleep(time.Second * 3)
	exitChan <- struct{}{}
	wg.Wait()
}
