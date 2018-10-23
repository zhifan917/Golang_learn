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

type Dog struct {
	name string
}

//Dog实现了上述2个接口的方法，所以其也实现了上述2个接口
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

	var a Animal
	fmt.Printf("a:%v dog:%v\n", a, dog)
	a = dog
	a.Eat()

	var b BuRuAnimal
	b = dog
	b.ChiNai()
}
