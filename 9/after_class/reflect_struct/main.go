package main

import (
	"fmt"
	"reflect"
)

type User struct {
	Name string `json:"name"`
	Age  int
	Sex  string
}

//1. 获取a的类型
//2. 我要动态改变a里面存的值
//3. 如果a里面存储的是一个结构体，那可以通过反射获取结构体中的字段信息以及调用结构体里面的方法
func TestValue(a interface{}) {

	v := reflect.ValueOf(a)
	t := v.Type()
	switch t.Kind() { //获取变量的类型
	case reflect.Struct: //假定变量是Struct结构体
		fieldNum := t.NumField() //通过NumField拿到结构体中的字段数量
		fmt.Printf("field num:%d\n", fieldNum)
		for i := 0; i < fieldNum; i++ {
			field := t.Field(i)  //字段的类型信息
			vField := v.Field(i) //变量的实例的值的相关信息

			fmt.Printf("field[%d] name:%s, json key:%s, val:%v\n",
				i, field.Name, field.Tag.Get("json"), vField.Interface()) //因为这里不确定值的类型，所以通过Interface()自动帮我们判别
		}
	}
}

func main() {

	var user User
	user.Name = "harden"
	user.Age = 100
	user.Sex = "man"
	TestValue(user)
	fmt.Printf("user:%#v\n", user)
}
