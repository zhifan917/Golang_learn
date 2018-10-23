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

func Update(Db *sqlx.DB) {
	username := "user02"
	age := 108
	id := 3
	result, err := Db.Exec("update user set name=?, age=? where id=?", username, age, id)
	if err != nil {
		fmt.Printf("exec failed, err:%v\n", err)
		return
	}

	affectRows, err := result.RowsAffected()
	if err != nil {
		fmt.Printf("affectRows failed, err:%v\n", err)
		return
	}

	fmt.Printf("last update id:%d affect rows:%d\n", id, affectRows)
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
	Update(Db)
}
