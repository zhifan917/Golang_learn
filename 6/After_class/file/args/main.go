package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Printf("args count:%d\n", len(os.Args)) //获取当前程序的参数个数
	for index, v := range os.Args {
		fmt.Printf("args%d, value:%s\n", index, v)
	}
}
