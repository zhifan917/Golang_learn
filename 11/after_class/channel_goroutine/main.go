package main

import (
	"fmt"
	"sync"
	"time"
)

var wg sync.WaitGroup

var exit bool

func worker(exitChan chan bool) {
LOOP: //设置一个标签
	for {
		fmt.Printf("worker\n")
		time.Sleep(time.Second)
		select {
		case <-exitChan:
			break LOOP //break标签LOOP 相当于退出标签下面对应的这个循环（注意：不是所有循环，而是紧接着的循环）
		default:

		}
	}
	wg.Done()

}

func main() {

	var exitChan chan bool = make(chan bool)
	wg.Add(1)
	go worker(exitChan)
	time.Sleep(time.Second * 3)
	exitChan <- true
	wg.Wait()
}
