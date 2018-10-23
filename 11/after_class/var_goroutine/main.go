package main

import (
	"fmt"
	"sync"
	"time"
)

var wg sync.WaitGroup

var exit bool

func worker(exitChan chan bool) {
	for {
		fmt.Printf("worker\n")
		time.Sleep(time.Second)
		//借助全局变量来退出
		if exit {
			break
		}
	}
	wg.Done()

}

func main() {

	var exitChan chan bool = make(chan bool)
	wg.Add(1)
	go worker(exitChan)
	time.Sleep(time.Second * 3)
	exit = true
	wg.Wait()
}
