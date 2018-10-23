package main

import (
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

type User struct {
	Id   int64  `db:"id"`
	Name string `db:"name"`
	Age  int    `db:"age"`
}

func Query(Db *sqlx.DB) {
	var user []*User //多行数据存在切片里
	id := 0
	err := Db.Select(&user, "select id, name, age from user where id>?", id)
	if err != nil {
		return
	}
	//fmt.Printf("user :%#v\n", user)
	for _, v := range user {
		fmt.Println(v)
	}
}

func main() {
	dns := "root:123456@tcp(127.0.0.1:3306)/golang"
	Db, err := sqlx.Connect("mysql", dns) //sqlx.connect连接数据库
	if err != nil {
		fmt.Printf("open mysql failed, err:%v\n", err)
		return
	}

	err = Db.Ping()
	if err != nil {
		fmt.Printf("ping failed, err:%v\n", err)
		return
	}

	fmt.Printf("connect to db succ\n")
	Query(Db)
}
