package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func Read(filename string) (string, error) {
	fi, err := os.Open(filename)
	if err != nil {
		return "", err
	}
	defer fi.Close()
	r := bufio.NewReader(fi)
	var content []byte
	var buf [1024]byte
	for {
		n, err := r.Read(buf[:])
		if err != nil && err != io.EOF {
			return "", err
		}
		if err == io.EOF {
			break
		}
		content = append(content, buf[:n]...)
	}
	fmt.Printf("content:%s\n", content)
	return string(content), nil
}

func main() {
	a := "c:/tmp.txt"
	Read(a)
}
