package main

import (
	"encoding/json"
	"fmt"
)

type Student struct {
	Name string
	Age  int
}

func (s *Student) String() string {
	data, _ := json.Marshal(s)
	return string(data)
}

func main() {
	var a = &Student{
		Name: "hell",
		Age:  12,
	}
	fmt.Printf("a = %v\n", a) //fmt包调Print相关函数时，看传进去的变量是否实现了stringer接口
}
