package main

import (
	"fmt"
	"os"
)

type Test struct {
	data string
}

//这里实现1个wirter接口的方法
func (t *Test) Write(p []byte) (n int, err error) {
	t.data = string(p)
	return len(p), nil
}

func main() {
	file, _ := os.Create("c:/tmp/c.txt")
	fmt.Fprintf(os.Stdout, "hello world\n") //输出到终端 （这里FPrintf函数要传入一个writer类型接口，把具体实例os.Stdout传给writer接口）
	fmt.Fprintf(file, "hello world\n")
	/* 因为writer接口的存在，我们可以省去下面的步骤：
	fmt.FPtrintfConsole()
	fmt.FPtrintfFile()
	fmt.FPtrintfNet()
	*/
	var t *Test = &Test{}
	fmt.Fprintf(t, "this is a test inteface:%s", "?akdfkdfjdkfk\n") //存入到data中了

	fmt.Printf("t.data:%s\n", t.data)
}
