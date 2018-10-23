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

	username := c.PostForm("username") //使用postform函数获取表单数据
	passwd := c.PostForm("passwd")

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
	r.POST("/user/info", handleUserInfo) //选择post方法

	r.Run(":9090")
}
