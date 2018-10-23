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

func testHSetGet(conn redis.Conn) {
	key := "abc"
	_, err := conn.Do("hset", "books", key, "this is a test") //books是哈希表名 其中存的是一条条key-value
	if err != nil {
		fmt.Printf("set failed:%s\n", err)
		return
	}

	//reply, err := conn.Do("get", "abc")
	data, err := redis.String(conn.Do("hget", "books", key))
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

	testHSetGet(conn)
}
