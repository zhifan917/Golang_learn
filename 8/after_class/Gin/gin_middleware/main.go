package main

import (
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func StatCost(c *gin.Context) { //统计耗时函数，这就是一个中间件
	start := time.Now()
	fmt.Printf("start stat cost\n")
	c.Next()                      ////等其他中间件先执行(也就是说执行完c.Next其他函数已经执行完了)
	lattancy := time.Since(start) //获取耗时
	fmt.Printf("process request cost:%d ms\n", lattancy/1000/1000)

}

func handleUserInfo(c *gin.Context) { //此请求处理函数就是在中间件函数的计算范围之内，干函数执行完，才会执行c.Next
	fmt.Printf("request start process\n")
	time.Sleep(3 * time.Second)
	c.JSON(http.StatusOK, "38333k333")
}

func main() {
	r := gin.Default()
	r.Use(StatCost) //将中间件函数设置（借助use函数）到gin框架，这样gin框架在处理请求时，就会先调用中间件进行处理

	r.GET("/user/info", handleUserInfo)
	r.Run(":8080")
}
