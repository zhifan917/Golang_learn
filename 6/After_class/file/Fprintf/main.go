package main

import (
	"fmt"
	"os"
)

func isFileExists(filename string) bool {
	_, err := os.Stat(filename)
	if os.IsNotExist(err) {
		return false
	}

	return true
}

func main() {
	filename := "c:/tmp.txt"

	var file *os.File
	var err error
	if isFileExists(filename) {
		//mac机器
		file, err = os.OpenFile(filename, os.O_APPEND|os.O_WRONLY, 0755)
	} else {
		file, err = os.Create(filename)
	}

	if err != nil {
		fmt.Printf("open %s failed, err:%v\n", filename, err)
		return
	}

	defer file.Close()

	fmt.Fprintf(file, "%d %d is good", 100, 300) //这里主要是为了演示一下Printf的底层调用Fprintf，我们可以通过传入文件对象，将其写入文件
}
