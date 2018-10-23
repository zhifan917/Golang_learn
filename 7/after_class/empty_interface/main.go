package main

import (
	"fmt"
)

func main() {
	var a interface{} //定义1个空接口
	var b int

	a = b //a是空接口，所以可以存储任何类型
	fmt.Printf("a=%v a:%T\n", a, a)
	var c float32

	a = c
	fmt.Printf("a=%v a:%T\n", a, a)
}
