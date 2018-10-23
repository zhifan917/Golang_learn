package main

import (
	"fmt"
	"reflect"
)

func TestType(a interface{}) {
	t := reflect.TypeOf(a) //获取a变量存的类型信息
	fmt.Printf("t = %v\n", t)

	kind := t.Kind()
	switch kind {
	case reflect.Int:
		fmt.Printf("a is int\n")
	case reflect.String:
		fmt.Printf("a is string\n")
	}
}

func TestValue(a interface{}) {
	v := reflect.ValueOf(a) //ValueOf函数返回的是一个value结构体，通过value就可以操作a，现在都赋值给v，所以我们就可以通过v去间接操作a，v也就相当于拿到了a的所有信息
	//注意v.Type() 和 reflect.TypeOf()的功能一样

	//动态改变a的值
	t := v.Type()
	switch t.Kind() {
	case reflect.Int:
		v.SetInt(1000)
	case reflect.String:
		v.SetString("xxxxxxx")
	case reflect.Ptr:
		t1 := v.Elem().Type() //v.Elem就相当于在变量前加个*号，获取该地址对应的值，此处就是获取指针变量对应的真实信息
		switch t1.Kind() {
		case reflect.Int:
			v.Elem().SetInt(1000)
			fmt.Printf("ptr is int\n")
		case reflect.String:
			v.Elem().SetString("xxxxxxx")
			fmt.Printf("ptr is string\n")
		}
		fmt.Printf("a is point type\n")
	}

}

func main() {
	var a int
	TestType(a)
	fmt.Printf("a=%v\n", a)
	var b string
	TestType(b)

	TestValue(&a)
	TestValue(&b)
	fmt.Printf("a=%v b=%v", a, b)
}
