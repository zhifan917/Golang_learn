package main

import (
	"fmt"
)

func describe(i interface{}) {
	fmt.Printf("Type = %T, value = %v\n", i, i)
}

func main() {
	s := "hello world"
	describe(s) //空接口可以存string
	i := 55
	describe(i) //空接口可以存int
	strt := struct {
		name string
	}{
		name: "Naveen R",
	}
	describe(strt) //空接口可以存struct，可以证明其可以存储任意类型
}
