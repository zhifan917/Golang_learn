package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func main() {
	filename := "c:/tmp.txt"
	file, err := os.Open(filename) //返回2个参数，第一个file指针，第二个错误值
	if err != nil {                //如果打开文件有错误，在这里输出
		fmt.Printf("open %s failed,err:%v\n", filename, err)
		return
	}
	defer func() { //打开一个文件，最后我们必须要将其关闭
		file.Close()
	}()

	reader := bufio.NewReader(file) //要封装这个文件读写，需要把文件传到bufio里面去，然后在bufio里面操作这个缓存,这里我们传的是file文件对象
	//为什么要传一个文件对象进去？   答：首先看缓存里有么有，有就读缓存，没有的话就用文件去读，也就是file这个文件对象（是一个结构体，里面包含很多用法）
	var content []byte //定义1个变量存读取到的所有数据
	var buf [4096]byte //定义一个4k的字节数组，每次读取一点，4k读性能高
	for {
		n, err := reader.Read(buf[:])    //将整个数组转换成切片读进去，Read函数返回2个参数，第1个n是读到的字节数，第二个是err
		if err != nil && err != io.EOF { //有一个特殊问题，当一个文件读读完，遇到文件末尾时，它也会返回一个错误，但是此时我已经读到文件末尾EOF，这个错误应该不算错误，所以应该把读到文件末尾这个错误给去掉。
			fmt.Printf("read %s failed, err:%v\n", filename, err)
			return //如果有错误就返回
		}

		if err == io.EOF {
			break //如果读取到文件末尾了，直接break退出。
		}

		vaildBuf := buf[0:n] //把有效数据都拿出来，不可能整个buf数组都是有效数据（最后一次读取到是很大可能是占据数组的一部分。），这里我们就需要借助切片。
		//fmt.Printf("%s\n", string(vaildBuf))
		content = append(content, vaildBuf...) //将有效的数据存到定义的变量切片中，另外将一个切片append到另一个切片中用...
	}
	fmt.Printf("content:%s\n", content)
}
