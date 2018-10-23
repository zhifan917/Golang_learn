package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type UserInfo struct {
	UserName string `form:"username" json:"username" binding:"required"` //form表单、json串
	Passwd   string `form:"passwd" json:"passwd" binding:"required"`
	Age      int    `form:"age" json:"age" binding:"required"`
	Sex      string `form:"sex" json:"sex" binding:"required"`
}

//form
func handleUserInfoForm(c *gin.Context) {
	var userInfo UserInfo
	err := c.ShouldBind(&userInfo) //直接将结构体的4个参数选项绑定过来，借助ShouldBind
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	} else {
		c.JSON(http.StatusOK, userInfo)
	}
}

//json
func handleUserInfoJson(c *gin.Context) {
	var userInfo UserInfo
	err := c.ShouldBindJSON(&userInfo)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	} else {
		c.JSON(http.StatusOK, userInfo)
	}
}

//querystring
func handleUserInfoQuery(c *gin.Context) {
	var userInfo UserInfo
	err := c.ShouldBind(&userInfo)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	} else {
		c.JSON(http.StatusOK, userInfo)
	}
}

func main() {
	r := gin.Default()

	// /v1/user/login
	// /v1/user/login2gi
	v1Group := r.Group("/v1")
	v1Group.POST("/user/infoform", handleUserInfoForm)
	v1Group.POST("/user/infojson", handleUserInfoJson)
	v1Group.GET("/user/infoquerystring", handleUserInfoQuery)
	r.Run(":9090")
}
