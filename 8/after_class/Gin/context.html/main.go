package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func handleHtml(c *gin.Context) {

	c.HTML(http.StatusOK, "post", "ksdkfskfkskfsdkfs") //post是html文件定义的名字
}

func main() {
	r := gin.Default()
	r.LoadHTMLGlob("./templates/*") //将模板文件加载进来

	r.GET("/user/info", handleHtml)

	r.Run(":9090")
}
