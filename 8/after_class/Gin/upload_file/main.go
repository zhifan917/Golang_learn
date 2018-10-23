package main

import (
	"fmt"
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

func handleUpload(c *gin.Context) {

	file, err := c.FormFile("testfile") //通过formfile方法来获取上传的文件，传入的参数是对上传文件的标识(可以随便起名),也就是要上传的文件的key
	if err != nil {
		fmt.Printf("upload file failed")
		return
	}

	filename := fmt.Sprintf("C:/tmp/%s", file.Filename) //文件保存位置
	err = c.SaveUploadedFile(file, filename)            //调用c将上传文件保存至指定位置
	if err != nil {
		fmt.Printf("save file failed, err:%v\n", err)
		return
	}

	c.JSON(http.StatusOK, "file upload success") //http.StatusOK就是返回状态码200
}

func main() {
	r := gin.Default() //Default返回一个默认的路由引擎
	r.POST("/file/upload", handleUpload)

	r.Run(":9090")
}
