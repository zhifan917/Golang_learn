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
	switch a.(type) { //格式是固定的，type是一个关键字  //强制转第一次type关键字
	case *Dog: //强制转换成dog（强制转第二次）
		dog := a.(*Dog)
		dog.Eat()
	case *Pig: //强制转成pig（强制转第二次）
		pig := a.(*Pig)
		pig.Eat()
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
