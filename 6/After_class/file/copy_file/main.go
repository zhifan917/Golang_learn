package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	CopyFile("d:/tmp.txt", "c:/tmp.txt") //目标文件，源文件
	fmt.Println("Copy done!")
}
func CopyFile(dstName, srcName string) (written int64, err error) {
	src, err := os.Open(srcName) //打开文件
	if err != nil {
		return
	}
	defer src.Close()
	dst, err := os.OpenFile(dstName, os.O_WRONLY|os.O_CREATE, 0644) //写入文件
	if err != nil {
		return
	}
	defer dst.Close()
	return io.Copy(dst, src) //核心在这里，io.Copy 一行代码搞定
}
