package main

import (
	"fmt"
	"io/ioutil"
)

func Read(filename string) (string, error) {
	content, err := ioutil.ReadFile(filename)
	if err != nil {
		return "", err
	}
	fmt.Printf("content:%s\n", content)
	return string(content), nil
}

func main() {
	a := "c:/tmp.txt"
	Read(a)
}
