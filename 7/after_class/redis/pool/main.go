package main

import (
	"fmt"
	"time"

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

func newPool(serverAddr string, passwd string) (pool *redis.Pool) { //第一个参数是服务的地址，第二个是连接时的密码（就是redis的密码）
	return &redis.Pool{ //返回1个结构体的连接池对象
		MaxIdle:     16,                //空闲连接数，即使没有连接请求，也会有空闲连接数在连接池中
		MaxActive:   1024,              //活跃连接数，也就是最大连接数
		IdleTimeout: 240 * time.Second, //空闲连接数超时时间，连接数超时了就会被释放掉
		Dial: func() (redis.Conn, error) { //匿名函数类型变量，用来连接redis
			fmt.Printf("create conn\n")
			conn, err := redis.Dial("tcp", serverAddr)
			if err != nil {
				return nil, err
			}

			if len(passwd) > 0 {
				_, err = conn.Do("auth", passwd)
				if err != nil {
					return nil, err
				}
			}
			return conn, err
		},
		TestOnBorrow: func(c redis.Conn, t time.Time) error { //匿名函数类型变量，作用：如果从连接池获取连接时，会验证一下这个连接是不是可用的。
			fmt.Printf("verify conn\n") //验证连接

			if time.Since(t) < time.Minute { //如果1分钟之内就不验证了，频繁的ping会影响性能
				return nil
			}
			fmt.Printf("ping conn\n")
			_, err := c.Do("ping")
			return err
		},
	}
}

func testRedisPool() {
	pool := newPool("127.0.0.1:6379", "")

	conn := pool.Get() //获取1个连接
	conn.Do("set", "abc", "3838383833834378473874837483748374")

	val, err := redis.String(conn.Do("get", "abc"))
	fmt.Printf("val:%s err:%v\n", val, err)

	//把连接归还到连接池，并不是关闭连接
	conn.Close()

	fmt.Printf("==========================\n")
	conn = pool.Get()
	conn.Do("set", "abc", "3838383833834378473874837483748374")

	val, err = redis.String(conn.Do("get", "abc"))
	fmt.Printf("val:%s err:%v\n", val, err)

	//把连接归还到连接池
	conn.Close()
}

func main() {
	testRedisPool()
}
