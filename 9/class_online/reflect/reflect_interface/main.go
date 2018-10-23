package main

import (
	"fmt"
	"reflect"
)

func TestType(a interface{}) {
	t := reflect.TypeOf(a)
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
	v := reflect.ValueOf(a)
	t := v.Type()
	switch t.Kind() {
	case reflect.Int:
		v.SetInt(1000)
	case reflect.String:
		v.SetString("xxxxxx")
	case reflect.Ptr:
		e := v.Elem()
		t1 := e.Type()
		switch t1.Kind() {
		case reflect.Int:
			v.Elem().SetInt(1000)
			fmt.Printf("ptr is int\n")
		case reflect.String:
			v.Elem().SetString("hello")
			fmt.Printf("ptr is string\n")
		}
		fmt.Printf("a is point type\n")
	}
}

func main() {
	var a int
	TestType(a)
	var b string
	TestType(b)
	fmt.Printf("a=%v\n", a)

	TestValue(&a)
	TestValue(&b)

	fmt.Printf("a=%v b=%v\n", a, b)
}
