package main

import (
	"fmt"
	"io/ioutil"
)

func Write() {
	fileName := "c:/tmp.txt"
	strTest := "测试测试"
	var d = []byte(strTest)
	err := ioutil.WriteFile(fileName, d, 0666)
	if err != nil {
		fmt.Println("write fail")
	}
	fmt.Println("write success")
}

func main() {
	Write()
}
