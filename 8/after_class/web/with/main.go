package main

import (
	"fmt"
	"html/template"
	"net/http"
)

var (
	t *template.Template
)

type Address struct {
	City     string
	Province string
	Code     string
}

type User struct {
	Name    string
	Age     int
	Address Address //继承
}

func initTemplate() (err error) {
	t, err = template.ParseFiles("./index.html")
	if err != nil {
		fmt.Printf("load template failed,err:%v\n", err)
		return
	}
	return
}

func handleUserInfo(w http.ResponseWriter, r *http.Request) {
	var user User = User{ //初始化
		Name: "user01",
		Age:  10,
		Address: Address{
			City:     "beijing",
			Province: "beijing",
			Code:     "10086",
		},
	}

	t.Execute(w, user)
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
