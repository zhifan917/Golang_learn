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

func handleMultiUpload(c *gin.Context) {

	form, err := c.MultipartForm() //多个文件用multipartform方法
	if err != nil {
		fmt.Printf("upload file failed")
		return
	}

	files := form.File["testfile"] //传入文件名的key，其返回值是一个数组
	for _, file := range files {   //遍历数组，即可把要传入的多个文件都保存起来了
		filename := fmt.Sprintf("C:/tmp/%s", file.Filename)
		err = c.SaveUploadedFile(file, filename)
		if err != nil {
			fmt.Printf("save file failed, err:%v\n", err)
			return
		}
	}
	c.JSON(http.StatusOK, "file upload success")
}

func main() {
	r := gin.Default()
	r.POST("/files/upload", handleMultiUpload)

	r.Run(":9090")
}
