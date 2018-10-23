package main

import (
	"fmt"
	"reflect"
)

func main() {
	var x float64 = 3.4
	//传地址进去，不传地址的话，改变的是副本的值
	//所以在reflect包里直接崩溃了!!!!!
	v := reflect.ValueOf(&x)
	fmt.Println("type:", v.Type())
	fmt.Println("kind is float64:", v.Kind() == reflect.Float64)
	fmt.Println("kind is point:", v.Kind() == reflect.Ptr)
	//通过Elem()获取指针指向的变量，从而完成赋值操作。
	//正常操作是通过*号来解决的，比如
	//var *p int = new(int)
	//*p = 100
	v.Elem().SetFloat(6.8)
	fmt.Println("value:", x)
}
