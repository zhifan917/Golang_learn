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

func PrepareQuery(DB *sql.DB) {
	//预处理第1步：发送sql语句命令部分
	stmt, err := DB.Prepare("select id, name, age from user where id>?")
	if err != nil {
		fmt.Printf("prepare failed, err:%v\n", err)
		return
	}

	//预处理第2步:发送sql语句数据部分
	id := 1
	rows, err := stmt.Query(id)
	//这个rows结果集，用完之后，一定释放！rows.Close()
	defer func() {
		if rows != nil {
			rows.Close()
		}
		if stmt != nil { //stmt记住也要释放
			stmt.Close()
		}
	}()
	if err == sql.ErrNoRows {
		fmt.Printf("not found data of id:%d\n", id)
		return
	}
	if err != nil {
		fmt.Printf("scan failed, err:%v\n", err)
		return
	}
	for rows.Next() {
		var user User
		err := rows.Scan(&user.Id, &user.Name, &user.Age)
		if err != nil {
			fmt.Printf("scan failed, err:%v\n", err)
			return
		}

		fmt.Printf("user:%#v\n", user)
	}
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

	PrepareQuery(Db)
}
