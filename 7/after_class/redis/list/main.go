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

func testList(conn redis.Conn) {

	_, err := conn.Do("lpush", "book_list", "this is a test", "ksksksksks") //左边进 第二个参数book_lsit是队列名
	if err != nil {
		fmt.Printf("set failed:%s\n", err)
		return
	}

	//reply, err := conn.Do("get", "abc")
	data, err := redis.String(conn.Do("rpop", "book_list")) //右边出，一次只能出队一个元素
	if err != nil {
		fmt.Printf("get failed, err:%v\n", err)
		return
	}

	fmt.Printf(" value:%s\n", data)
}

func main() {
	conn, err := initRedis()
	if err != nil {
		return
	}
	defer conn.Close() //关闭连接

	testList(conn)
}
