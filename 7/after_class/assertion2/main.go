package main

import (
	"fmt"
)

func assert(i interface{}) {
	v, ok := i.(int) //将空接口传入的类型定为int
	fmt.Println(v, ok)
}

func main() {
	var s interface{} = 56
	assert(s)
	var i interface{} = "harden"
	assert(i)
}
