package main

import (
	"fmt"
	"math/rand"
	"net/http"
	"time"
)

func indexHandle(w http.ResponseWriter, r *http.Request) {
	number := rand.Intn(2)
	if number == 0 {
		time.Sleep(time.Second * 10)
		fmt.Fprintf(w, "slow request")
		return
	}
	fmt.Fprintf(w, "quick response")
}

func main() {
	http.HandleFunc("/", indexHandle)
	http.ListenAndServe(":10000", nil)
}
