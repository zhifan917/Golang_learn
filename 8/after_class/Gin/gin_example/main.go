package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Result struct {
	Message string `json:"message"`
	Code    int    `json:"code"`
}

func handleUserInfo(c *gin.Context) {

	var result Result = Result{ //初始化1个map
		Message: "success",
		Code:    0,
	}

	c.JSON(http.StatusOK, result) //第二个参数是一个空接口，可以传递任何类型数据进来，这里传递进来的map会自动把我们转成json串
}

func main() {
	r := gin.Default()
	r.GET("/user/info", handleUserInfo)

	r.Run(":9090")
}
