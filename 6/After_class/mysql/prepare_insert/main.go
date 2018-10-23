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

func PrepareInsert(DB *sql.DB) {
	stmt, err := DB.Prepare("insert into user(name, age)values(?, ?)")
	if err != nil {
		fmt.Printf("prepare failed, err:%v\n", err)
		return
	}

	result, err := stmt.Exec("user10111", 108)
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

	PrepareInsert(Db)
}
