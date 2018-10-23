package main

import (
	"fmt"
	"log"
	"net/http"
	"strings"
)

func sayhelloName(w http.ResponseWriter, r *http.Request) { //第一个参数是一个接口，第二个参数是一个结构体（封装了一系列方法，比如：取post提交参数等），r是接受请求，w是返回给浏览器
	r.ParseForm()                       //解析参数，默认是不会解析的，如果是post请求（带表单参数），会进行解析
	fmt.Println(r.Form)                 //这些信息是输出到服务器端的打印信息，将表单打印出来
	fmt.Println("path", r.URL.Path)     //请求路径
	fmt.Println("scheme", r.URL.Scheme) //scheme是http还是https
	fmt.Println(r.Form["url_long"])     //传递了一个表单参数url_long
	for k, v := range r.Form {
		fmt.Println("key:", k)
		fmt.Println("val:", strings.Join(v, ""))
	}
	fmt.Fprintf(w, "Hello world!") //这个写入到w的是输出到客户端的
	//w这里是一个接口，接口的好处就体现在封装成一个接口的话，不论是返回文件或者图片或者html都可以通过接口进行返回，不用接口的话，你需要为每一种返回类型写函数。帮你更好理解

}
func main() {
	http.HandleFunc("/", sayhelloName)       //设置访问的路由，sayhelloName是函数类型的参数
	err := http.ListenAndServe(":9090", nil) //设置监听的端口
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
