package main

import (
	"bufio"
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

	file, err = os.Create(filename)
	if err != nil {
		fmt.Printf("open %s failed, err:%v\n", filename, err)
		return
	}

	defer file.Close()

	writer := bufio.NewWriter(file)
	writer.WriteString("hello worldldfdsfsfsf")

	writer.Flush()
}
