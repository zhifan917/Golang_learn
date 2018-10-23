package main

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	dsn := "root:123456@tcp(127.0.0.1:3306)/golang" //连接数据库的配置
	Db, err := sql.Open("mysql", dsn)               //Open函数第一个参数就是驱动的名字（不能随意写的）
	if err != nil {                                 //有一个坑，如果连接数据库配置写错了，不会在此处报错
		fmt.Printf("open mysql is failed,err:%v\n", err)
		return
	}

	err = Db.Ping() //ping一下没报错证明连接数据库成功
	if err != nil {
		fmt.Printf("ping is failed,err:%v\n", err)
		return
	}

	fmt.Printf("connect to db successful\n")
}
