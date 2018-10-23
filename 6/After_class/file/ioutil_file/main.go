package main

import (
	"fmt"
	"io/ioutil"
)

func main() {
	filename := "c:/tmp.txt"
	str := "dkfslfjdsklfjlskjflsjflsjflsjflks"
	err := ioutil.WriteFile(filename, []byte(str), 0755)
	if err != nil {
		fmt.Println("write fail")
	}
	fmt.Println("write success")
}
