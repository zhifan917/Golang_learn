package main

import (
	"fmt"

	"github.com/garyburd/redigo/redis"
)

func initRedis() (conn redis.Conn, err error) { //连接redis函数
	conn, err = redis.Dial("tcp", "127.0.0.1:6379") //传递协议、ip、端口
	if err != nil {
		fmt.Printf("conn redis failed, err:%v\n", err)
		return
	}
	fmt.Printf("connect redis successful!!!\n")
	return
}

func testSetGet(conn redis.Conn) {
	key := "abc"
	_, err := conn.Do("set", key, "this is a test") //用do函数来进行redis命令操作
	if err != nil {
		fmt.Printf("set failed:%s\n", err)
		return
	}

	//reply, err := conn.Do("get", "abc") //get返回的是1个空接口，我们不知道里面内容到底什么类型，所以要做一次转换
	data, err := redis.String(conn.Do("get", key)) //因为我们知道存的是string，所以转换时是redis.string，如果存的是int，那就是redis.int
	if err != nil {
		fmt.Printf("get failed, err:%v\n", err)
		return
	}

	fmt.Printf("key:%s value:%s\n", key, data)
}

func main() {
	conn, err := initRedis()
	if err != nil {
		return
	}
	defer conn.Close() //关闭连接

	testSetGet(conn)
}
