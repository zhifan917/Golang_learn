package main

import (
	"fmt"
)

//定义2个接口
type Animal interface {
	Eat()
}

type BuRuAnimal interface {
	ChiNai()
}

//接口嵌套
type AdvanceAnimal interface { //要想实现AdvanceAnimal接口，那么就需要满足嵌套的接口的所有方法
	Animal
	BuRuAnimal
}

type Dog struct {
	name string
}

func (d *Dog) Eat() {
	fmt.Printf("%s is eating\n", d.name)
}

func (d *Dog) ChiNai() {
	fmt.Printf("%s is ChiNai\n", d.name)
}

func main() {
	var dog = &Dog{
		name: "旺财",
	}

	var a AdvanceAnimal
	fmt.Printf("a:%v dog:%v\n", a, dog)
	a = dog
	a.Eat()
	a.ChiNai()
}
