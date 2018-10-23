package main

import (
	"fmt"
)

func findType(i interface{}) {
	switch v := i.(type) {
	case string:
		fmt.Printf("I am a string and my value is %s\n", v)
	case int:
		fmt.Printf("I am a int and my value is %d\n", v)
	default:
		fmt.Printf("Unknown type\n")
	}
}

func main() {
	findType("hello")
	findType(77)
	findType(88.98)
}
