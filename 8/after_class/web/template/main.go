package main

import (
	"fmt"
	"html/template" //加载模板的包
	"net/http"
)

var (
	t *template.Template //定义一个接收模板的全局变量
)

type User struct {
	Name string
	Age  int
}

func initTemplate() (err error) {
	t, err = template.ParseFiles("./index.html") //把模板加载进来 ，t是返回的实例
	if err != nil {
		fmt.Printf("load template failed,err:%v\n", err)
		return
	}
	return
}

func handleUserInfo(w http.ResponseWriter, r *http.Request) {
	var user User = User{ //初始化  //可以传递任何类型的数据进去
		Name: "user01",
		Age:  10,
	}

	t.Execute(w, user) //进行渲染，t对应的就是模板中的.
}

func main() {

	err := initTemplate()
	if err != nil {
		return
	}

	http.HandleFunc("/user/info", handleUserInfo)
	err = http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Printf("listen failed, err:%v\n", err)
		return
	}
}
