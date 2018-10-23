package main

import (
	"flag"
	"fmt"
)

func main() {
	var num int
	var mode string

	flag.IntVar(&num, "num", 16, "-num the passwd length") //IntVar参数依次是：传入值-终端参数-默认值-用法
	flag.StringVar(&mode, "mode", "mix", "-mode the password generate mode")

	flag.Parse() //真正解析命令行参数，IntVar、StringVar只是设置命令行需要的一些参数

	fmt.Printf("num:%d mode:%s\n", num, mode)
}
