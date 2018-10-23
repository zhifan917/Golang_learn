package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

func greet(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "greet hello World! %s", time.Now())
}

func userLogin(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		data, err := ioutil.ReadFile("./login.html")
		if err != nil {
			http.Redirect(w, r, "404.html", http.StatusFound)
			return
		}
		w.Write(data)
	} else if r.Method == "POST" {
		r.ParseForm()
		username := r.FormValue("username")
		password := r.FormValue("password")

		if username == "admin" && password == "admin" {
			fmt.Fprintf(w, "login success")
		} else {
			fmt.Fprintf(w, "login failed")
		}
	}
}

func main() {
	http.HandleFunc("/", greet)
	http.HandleFunc("/user/login", userLogin)
	http.ListenAndServe(":8080", nil)
}
