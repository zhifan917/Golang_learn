package main

import (
	"fmt"
	"os"
)

func isFileExists(filename string) bool { //判断文件是否存在
	_, err := os.Stat(filename) //os.Stat会返回文件是否存在的相关属性
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
		file, err = os.OpenFile(filename, os.O_APPEND, 0755) //如果文件存在则追加进去
		//file, err = os.OpenFile(filename, os.O_APPEND|os.O_WRONLY, 0755) //mac系统追加时还需要在加一个os.O_WRONLY方法
	} else {
		file, err = os.Create(filename) //不存在就创建该文件
	}

	if err != nil {
		fmt.Printf("open %s failed, err:%v\n", filename, err)
		return
	}

	defer file.Close()

	//给文件中写入内容
	n, err := file.WriteString("hello world") //WriteString可以传入字符串
	//io.WriteString(file,"hello world")   //io.WriteString也可以用来进行传入
	if err != nil {
		fmt.Printf("write failed, err:%v\n", err)
		return
	}
	fmt.Printf("write %d succ\n", n)
}
