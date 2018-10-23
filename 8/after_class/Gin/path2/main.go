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

func handleUserParams(c *gin.Context) {

	username := c.Param("username")
	passwd := c.Param("passwd")

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
	r.GET("/user/info/:username/:passwd", handleUserParams)

	r.Run(":9090")
}
