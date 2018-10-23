package main

import (
	"fmt"
	"sync"

	"github.com/Shopify/sarama"
)

var wg sync.WaitGroup

func main() {
	consumer, err := sarama.NewConsumer([]string{"192.168.20.200:9092"}, nil)
	if err != nil {
		fmt.Println("consumer close,err:", err)
		return
	}

	fmt.Printf("connect succ\n")
	partions, err := consumer.Partitions("zzf_log")
	if err != nil {
		fmt.Printf("get partition failed, err:%v\n", err)
		return
	}

	for _, p := range partions {
		pc, err := consumer.ConsumePartition("zzf_log", p, sarama.OffsetNewest)
		if err != nil {
			fmt.Printf("err:%v\n", err)
			continue
		}
		wg.Add(1)
		go func() {
			for m := range pc.Messages() {
				fmt.Printf("message:%v, text:%v\n", m, string(m.Value))
			}
			wg.Done()
		}()
	}
	wg.Wait()
}
