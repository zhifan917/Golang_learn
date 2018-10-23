package main

import (
	"fmt"
)

type Animal interface {
	Eat()
	Talk()
	Run()
}

type Dog struct {
	name string
}

func (d *Dog) Eat() {
	fmt.Printf("%s is eating\n", d.name)
}

func (d *Dog) Talk() {
	fmt.Printf("%s is talking\n", d.name)
}

func (d *Dog) Run() {
	fmt.Printf("%s is running\n", d.name)
}

func main() {
	var dog = &Dog{
		name: "旺财",
	}

	var a Animal
	fmt.Printf("a:%v dog:%v\n", a, dog)

	a = dog
	a.Eat()
	a.Run()
	a.Talk()
}
