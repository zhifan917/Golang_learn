package main

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

type User struct {
	Id   int64  `db:"id"`
	Name string `db:"name"`
	Age  int    `db:"age"`
}

func QueryRow(Db *sqlx.DB) { //查询单行数据
	id := 2
	var user User
	err := Db.Get(&user, "select id, name, age from user where id=?", id)
	if err == sql.ErrNoRows { //无数据报错
		fmt.Printf("not record found\n")
		return
	}
	if err != nil { //查询失败报错
		fmt.Printf("get failed, err:%v\n", err)
		return
	}

	fmt.Printf("get user succ, user:%#v\n", user)
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
	QueryRow(Db)
}
