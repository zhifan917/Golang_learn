package reflect_json

import (
	"fmt"
	"reflect"
)

func Marshal(data interface{}) (jsonStr string) { //data就是用户传进来的数据信息，返回的是序列化完的json序列串

	t := reflect.TypeOf(data) //获取类型信息
	v := reflect.ValueOf(data)
	switch t.Kind() { //猜是什么类型
	case reflect.String, reflect.Int, reflect.Int32: //简单数据类型处理几乎是一致的
		jsonStr = fmt.Sprintf("\"%v\"", data) // \"表示双引号
	case reflect.Struct: //结构体
		numField := t.NumField() //获取结构体字段数量
		for i := 0; i < numField; i++ {
			//类型信息
			name := t.Field(i).Name           //获取字段名字
			tag := t.Field(i).Tag.Get("json") //获取有tag的字段名（此处针对json）
			if len(tag) > 0 {                 //有tag，优先使用tag
				name = tag
			}
			//值信息
			vField := v.Field(i)              //返回值是一个Value的结构体
			vFieldValue := vField.Interface() //想要获取字段的值，用interface即可
			//拼接json
			if t.Field(i).Type.Kind() == reflect.String {
				jsonStr += fmt.Sprintf("\"%s\":\"%v\"", name, vFieldValue) //如果是字符串就加双引号
			} else {
				jsonStr += fmt.Sprintf("\"%s\":%v", name, vFieldValue) //不是字符串值不用加双引号
			}

			//json串的话，字段与字段之间还有1个单引号，最后一个字段没有逗号
			if i != numField-1 { //如果不是最后一个，就加一个逗号
				jsonStr += ","
			}
		}

		jsonStr = "{" + jsonStr + "}" //最后在最前面和最后面加一个大括号
	}
	return
}
