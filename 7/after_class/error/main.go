package main

import (
	"fmt"
	"time"
)

type MyError struct {
	When time.Time
	What string
}

func (e MyError) Error() string {
	str := fmt.Sprintf("at %v, %s", e.When, e.What)
	fmt.Printf("1:%T\n", str)
	return str
}
func run() error {
	fmt.Println("0")
	str := MyError{time.Now(), "it didn't work"}
	fmt.Printf("2:%T\n", str)
	fmt.Println(MyError{time.Now(), "it didn't work"})
	return str
}
func main() {
	if err := run(); err != nil {
		fmt.Printf("3:%T\n", err)
		fmt.Println(err)
	}
}
