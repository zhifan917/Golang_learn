package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	//标准输出
	writer := bufio.NewWriter(os.Stdout) //写的话相当于从终端输出，读的话相当于从终端输入
	writer.WriteString("hello world")
	writer.Flush() //因为bufio还写在内存中，需要flush一下

	fmt.Printf("hello world")
	//标准输入
	reader := bufio.NewReader(os.Stdin)  //从终端读入
	data, err := reader.ReadString('\n') //readstring方法是传入分隔符，此处传入\n表示直到读入\n为止，也就相当于读一行内容即可
	if err != nil {
		fmt.Printf("read from console failed, err:%v\n", err)
		return
	}
	fmt.Fprintf(os.Stdout, "data:%s\n", data)
}
