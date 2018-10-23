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

func testExpire(conn redis.Conn) {
	_, err := conn.Do("expire", "abc", 20)
	if err != nil {
		fmt.Println(err)
		return
	}
}

func main() {
	conn, err := initRedis()
	if err != nil {
		return
	}
	defer conn.Close() //关闭连接

	testExpire(conn)
}
