package main

import (
	"fmt"
)

type Animal interface {
	Eat()
}

type Dog struct {
	name string
}

func (d *Dog) Eat() {
	fmt.Printf("%s is eating\n", d.name)
}

type Pig struct {
	name string
}

func (d *Pig) Eat() {
	fmt.Printf("%s is eating\n", d.name)
}

//types switch
func DescribeSwitch(a Animal) {
	fmt.Printf("DescribeSwitch(a) begin\n")
	switch v := a.(type) { //格式是固定的，type是一个关键字
	case *Dog:
		v.Eat() //v就是断言之后的具体类型
	case *Pig:
		v.Eat()
	}
	fmt.Printf("DescribeSwitch(a) end\n")
}

func main() {
	var dog = &Dog{
		name: "旺财",
	}

	var a Animal
	a = dog
	fmt.Printf("I am dog\n")
	DescribeSwitch(a)

	var pig = &Pig{
		name: "佩奇",
	}

	a = pig
	fmt.Printf("I am pig\n")
	DescribeSwitch(a)
}
