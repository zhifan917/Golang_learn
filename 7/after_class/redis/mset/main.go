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

func testMSetGet(conn redis.Conn) {
	key := "abc"
	key1 := "efg"
	_, err := conn.Do("mset", key, "this is a test", key1, "ksksksksks") //一次设置多个key
	if err != nil {
		fmt.Printf("set failed:%s\n", err)
		return
	}

	//reply, err := conn.Do("get", "abc")
	data, err := redis.Strings(conn.Do("mget", key, key1)) //一次读取多个 。返回的data是1个切片，这里用strings函数
	if err != nil {
		fmt.Printf("get failed, err:%v\n", err)
		return
	}
	for _, val := range data { //遍历出来
		fmt.Printf(" value:%s\n", val)
	}
}

func main() {
	conn, err := initRedis()
	if err != nil {
		return
	}
	defer conn.Close() //关闭连接

	testMSetGet(conn)
}
