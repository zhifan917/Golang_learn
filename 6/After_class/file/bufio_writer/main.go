package main

import (
	"bufio"
	"fmt"
	"os"
)

func Write() {
	fileName := "c:/tmp.txt"
	f, err3 := os.Create(fileName) //创建文件
	if err3 != nil {
		fmt.Println("create file fail")
	}
	w := bufio.NewWriter(f)                //创建新的 Writer 对象
	n4, err3 := w.WriteString("bufferedn") //此时还是写入在内存中
	fmt.Printf("写入 %d 个字节\n", n4)
	w.Flush() //刷新了磁盘
	defer f.Close()
}

func main() {
	Write()
}
