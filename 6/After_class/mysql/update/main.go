package main

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

type User struct {
	Id   int64  `db:"id"`
	Name string `db:"name"`
	Age  int    `db:"age"`
}

func Update(DB *sql.DB) {
	username := "user02"
	age := 108
	id := 3
	result, err := DB.Exec("update user set name=?, age=? where id=?", username, age, id)
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
	dsn := "root:123456@tcp(127.0.0.1:3306)/golang"
	Db, err := sql.Open("mysql", dsn)
	if err != nil {
		fmt.Printf("open mysql is failed,err:%v\n", err)
		return
	}

	err = Db.Ping()
	if err != nil {
		fmt.Printf("ping is failed,err:%v\n", err)
		return
	}

	fmt.Printf("connect to db successful\n")

	Update(Db)
}
