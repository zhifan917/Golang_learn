package main

import "github.com/gin-gonic/gin"

func main() {
	r := gin.Default()                    //Default返回一个默认的路由引擎
	r.GET("/ping", func(c *gin.Context) { //c这里就是封装了go底层web的r和w
		c.JSON(200, gin.H{ //200是状态码
			"message": "pong",
		}) //输出json结果给调用方
	})
	r.Run() // listen and serve on 0.0.0.0:8080
}
