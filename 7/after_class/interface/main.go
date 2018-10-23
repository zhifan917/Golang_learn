package main

import (
	"fmt"
)

type Animal interface { //定义了动物的规范（接口定义的一组方法）
	Eat()
	Talk()
	Run()
}

type Dog struct { //狗如果能够满足了动物的规范（接口方法），那其就是动物
	name string
}

func (d *Dog) Eat() { //目前狗还不是动物，因为其只实现了Eat，还需要实现Talk和Run，其才算是动物
	fmt.Printf("%s is eating\n", d.name)
}

func (d *Dog) Talk() {
	fmt.Printf("%s is talking\n", d.name)
}

func (d *Dog) Run() {
	fmt.Printf("%s is runing\n", d.name)
}

type Pig struct {
	name string
}

func (d *Pig) Eat() {
	fmt.Printf("%s is eating\n", d.name)
}

func (d *Pig) Talk() {
	fmt.Printf("%s is talking\n", d.name)
}

func (d *Pig) Run() {
	fmt.Printf("%s is runing\n", d.name)
}

func main() {
	var dog = &Dog{
		name: "旺财",
	}

	var a Animal
	fmt.Printf("a:%v dog:%v\n", a, dog) //接口底层就是指针，指向的就是一个空对象，如果字节调用就会panic

	a = dog //dog满足了接口所有方法，所以我们可以直接将其复制给Animal，对应的理论就是接口类型的变量可以保存实现该接口的任何具体类型的实例。
	a.Eat()
	a.Run()
	a.Talk()

	var pig = &Pig{
		name: "佩奇",
	}

	a = pig
	a.Eat()
	a.Run()
	a.Talk()
}
