package main

import (
	"fmt"
	_ "strings"
	"time"
)

type MyReader struct{}

// TODO: Add a Read([]byte) (int, error) method to MyReader.
func (r MyReader) Read(b []byte) (int, error) {
	b[0] = 'A'
	return 1, nil
}
func main() {
	var myre MyReader
	b := make([]byte, 1)
	//for {
	//r := strings.NewReader(b)
	myre.Read(b)
	fmt.Printf("%c\n", b[0])
	time.Sleep(1 * time.Second)
	myre.Read(b)
	fmt.Println(b[0])
	//}
}
