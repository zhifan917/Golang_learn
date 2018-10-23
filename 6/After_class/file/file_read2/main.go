package main

import (
	"fmt"
	"io"
	"os"
)

func Read(filename string) (string, error) {
	//获得一个file
	f, err := os.Open(filename)
	if err != nil {
		fmt.Println("read fail")
		return "", err
	}

	//把file读取到缓冲区中
	defer f.Close()
	var content []byte
	var buf [1024]byte
	for {
		//从file读取到buf中， n表示本次读到的字节数
		n, err := f.Read(buf[:])
		if err != nil && err != io.EOF {
			fmt.Println("read buf fail", err)
			return "", err
		}
		//说明读取结束
		if err == io.EOF {
			break
		}
		//读取到最终的缓冲区中
		content = append(content, buf[:n]...)
	}
	fmt.Printf("content:%s\n", content)
	return string(content), nil
}

func main() {
	a := "c:/tmp.txt"
	Read(a)
}
