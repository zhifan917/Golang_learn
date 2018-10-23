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

//类型断言
func Describe(a Animal) { //注意Animal（接口）本身就是引用类型，这里不能加*，加上会报错
	dog, ok := a.(*Dog) //引入ok判断机制
	if !ok {
		fmt.Printf("convert to dog failed\n")
		return
	}
	fmt.Printf("describe suncc\n")
	dog.Eat()
	fmt.Printf("describe suncc----------\n")
}

func main() {
	var dog = &Dog{
		name: "旺财",
	}

	var a Animal
	a = dog
	fmt.Printf("I am dog\n")
	Describe(a)

	var pig = &Pig{
		name: "佩奇",
	}

	a = pig
	fmt.Printf("I am pig\n")
	Describe(a)
}
