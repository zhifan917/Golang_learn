package main

import (
	"fmt"
	"io/ioutil"
)

func main() {
	filename := "c:/tmp.txt"
	content, err := ioutil.ReadFile(filename)
	if err != nil {
		fmt.Printf("read file %s failed, err:%v\n", filename, err)
		return
	}
	fmt.Printf("content:%s\n", string(content)) //因为上面返回的是一个字节数组，所以必须转一下
}
