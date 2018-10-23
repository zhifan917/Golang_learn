package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func userLogin(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("r.method:%s\n", r.Method)
	if r.Method == "GET" { //如果是get直接返回请求页面
		data, err := ioutil.ReadFile("./login.html")
		if err != nil {
			http.Redirect(w, r, "/404.html", http.StatusNotFound)
			return
		}

		w.Write(data) //要把内容输出到浏览器
	} else if r.Method == "POST" {
		r.ParseForm()                       //解析表单
		username := r.FormValue("username") //和html文件要一一对应
		password := r.FormValue("password")

		if username == "admin" && password == "admin" {
			fmt.Fprintf(w, "login success")
		} else {
			fmt.Fprintf(w, "login failed")
		}
	}
}

func main() {
	http.HandleFunc("/user/login", userLogin)
	err := http.ListenAndServe(":9090", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
