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

func Insert(DB *sql.DB) {
	username := "user01"
	age := 18

	result, err := DB.Exec("insert into user(name, age)values(?, ?)", username, age)
	if err != nil {
		fmt.Printf("exec failed, err:%v\n", err)
		return
	}

	id, err := result.LastInsertId() //获取最后一次插入的id
	if err != nil {
		fmt.Printf("last insert id failed, err:%v\n", err)
		return
	}

	affectRows, err := result.RowsAffected() //此次操作影响的行数
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

	Insert(Db)
}
