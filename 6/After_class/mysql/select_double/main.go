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

func Query(Db *sql.DB) {
	id := 0
	rows, err := Db.Query("select id, name, age from user where id>?", id)
	//rows返回的是1个结果集，这个rows结果集，用完之后，一定释放！rows.Close() 这个是消耗数据库连接资源的。
	defer func() {
		if rows != nil {
			rows.Close()
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

	for rows.Next() { //多行数据通过写一个for循环来进行处理，有数据返回true接着执行，没数据返回false就退出
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

	Query(Db)
}
