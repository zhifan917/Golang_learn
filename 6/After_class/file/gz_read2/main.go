package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func main() {
	fmt.Printf("start run...\n")

	filename := "c:/tmp.txt.gz"
	file, err := os.Open(filename)
	if err != nil {
		fmt.Printf("open %s failed, err:%v\n", filename, err)
		return
	}
	fmt.Printf("start0 read file\n")
	defer file.Close()

	reader := bufio.NewReader(file)
	var content []byte
	var buf [4096]byte
	for {
		n, err := reader.Read(buf[:])
		if err != nil && err != io.EOF {
			fmt.Printf("read %s failed, err:%v\n", filename, err)
			return
		}

		if err == io.EOF {
			break
		}

		vaildBuf := buf[0:n]
		content = append(content, vaildBuf...)
	}
	fmt.Printf("content:%s\n", content)
}
