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

func Insert(Db *sqlx.DB) {
	username := "harden"
	age := 18

	result, err := Db.Exec("insert into user(name, age)values(?, ?)", username, age)
	if err != nil {
		fmt.Printf("exec failed, err:%v\n", err)
		return
	}

	id, err := result.LastInsertId()
	if err != nil {
		fmt.Printf("last insert id failed, err:%v\n", err)
		return
	}

	affectRows, err := result.RowsAffected()
	if err != nil {
		fmt.Printf("affectRows failed, err:%v\n", err)
		return
	}

	fmt.Printf("last insert id:%d affect rows:%d\n", id, affectRows)
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
	Insert(Db)
}
