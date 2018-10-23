package main

import (
	"fmt"
	"net/http"
)

func greet(w http.ResponseWriter, r *http.Request) {
	word := r.FormValue("word") //通过formvalue来获取用户传递过来参数
	fmt.Fprintf(w, "greet Hello World! word:%s", word)
}

func main() {
	http.HandleFunc("/index", greet)
	http.ListenAndServe(":8080", nil)
}
