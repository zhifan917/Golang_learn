package main

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	dsn := "root:123456@tcp(192.168.20.200:3306)/golang"
	//dsn := "root:@tcp(127.0.0.1:3306)/test"
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		fmt.Printf("open mysql failed,err:%v\n", err)
		return
	}

	err = db.Ping()
	if err != nil {
		fmt.Printf("ping mysql failed,err:%v\n", err)
		return
	}
	fmt.Printf("connect to db succ\n")
}
