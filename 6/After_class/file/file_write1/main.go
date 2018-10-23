package main

import (
	"io"
	"os"
)

func CheckFileExist(fileName string) bool {
	_, err := os.Stat(fileName)
	if os.IsNotExist(err) {
		return false
	}
	return true
}

func Write(fileName string) error {
	var f *os.File
	var err error
	if CheckFileExist(fileName) { //文件存在
		f, err = os.OpenFile(fileName, os.O_APPEND, 0666) //打开文件
		if err != nil {
			return err
		}
	} else { //文件不存在
		f, err = os.Create(fileName) //创建文件
		if err != nil {
			return err
		}
	}
	_, err = io.WriteString(f, "strTest")
	if err != nil {
		return err
	}
	return nil
}

func main() {
	Write("c:/tmp.txt")
}
