package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Result struct {
	Message string `json:"message"`
	Code    int    `json:"code"`
}

type UserInfo struct {
	Result
	UserName string `json:"username"`
	Passwd   string `json:"passwd"`
}

func handleUserInfo(c *gin.Context) {

	username := c.Query("username")
	passwd := c.DefaultQuery("passwd", "dkdkdkdkdkdkdkd") //默认值，如果用户没有传递过来值，则为此处定义默认值

	var result UserInfo = UserInfo{
		UserName: username,
		Passwd:   passwd,
	}

	result.Code = 0
	result.Message = "success"

	c.JSON(http.StatusOK, result)
}

func main() {
	r := gin.Default()
	r.GET("/user/info", handleUserInfo)

	r.Run(":9090")
}
