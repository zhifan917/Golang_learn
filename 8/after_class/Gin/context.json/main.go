package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	// gin.H is a shortcut for map[string]interface{}
	r.GET("/someJSON", func(c *gin.Context) {
		//第一种方式,自己拼json
		c.JSON(http.StatusOK, gin.H{"message": "hey", "status": http.StatusOK}) //gin.H就是一个map的别名，借助gin.H可以去拼接json
	})

	//第二种：通过定义一个结构体，然后将结构体传给c.JSON，之后会将传递过来的结构体打包成支持json协议的字符串返回给客户端
	r.GET("/moreJSON", func(c *gin.Context) {
		// You also can use a struct
		var msg struct {
			Name    string `json:"user"`
			Message string
			Number  int
		}
		msg.Name = "Lena"
		msg.Message = "hey"
		msg.Number = 123
		// Note that msg.Name becomes "user" in the JSON
		c.JSON(http.StatusOK, msg)
	})
	// Listen and serve on 0.0.0.0:8080
	r.Run(":8080")
}
