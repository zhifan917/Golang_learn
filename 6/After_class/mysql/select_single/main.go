package main

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

type User struct { //为要查询的表定义一个结构体（数据库定义的表要有1个结构体来封装，否则会十分混乱）
	Id   int64  `db:"id"` //数据库一般都是小写，所以我们需要借助tag来映射(结构体中字段名写什么就无所谓了，反正要做映射)
	Name string `db:"name"`
	Age  int    `db:"age"`
}

func QueryRow(Db *sql.DB) {
	id := 1
	row := Db.QueryRow("select id,name,age from user where id=?", id) //?是占位符，可以传入数值拼接成完整的sql语句发给mysql服务器，然后mysql返回数据集。

	var user User
	err := row.Scan(&user.Id, &user.Name, &user.Age) //Scan这里查询有2种错误 1、查询不到数据失败（网络原因等）；2、查询过程中查询失败（和因为网络原因造成的错误区分开）
	if err == sql.ErrNoRows {                        //对应第二种错误 ErrNoRows表示查询过程中是没发生错误的
		fmt.Printf("not found data of id:%d\n", id)
		return
	}

	if err != nil { //对应第一种错误
		fmt.Printf("scan failed, err:%v\n", err)
		return
	}

	fmt.Printf("user:%#v\n", user)

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

	QueryRow(Db)
}
