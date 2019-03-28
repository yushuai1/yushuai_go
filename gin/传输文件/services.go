package main

import (
	"github.com/gin-gonic/gin"
	"os"
	"log"
	"fmt"
	"io"
	"net/http"
)

// 接收client上传的文件
// 从FormFile中获取相关的文件data!
// 然后写入本地文件
func func9(c *gin.Context) {
	// 注意此处的文件名和client处的应该是一样的
	file, header, err := c.Request.FormFile("uploadFile")
	filename := header.Filename
	fmt.Println(header.Filename)
	// 创建临时接收文件
	out, err := os.Create("copy_" + filename)
	if err != nil {
		log.Fatal(err)
	}
	defer out.Close()
	// Copy数据
	_, err = io.Copy(out, file)
	if err != nil {
		log.Fatal(err)
	}
	c.String(http.StatusOK, "upload file success")
}

func main() {
	router := gin.Default()
	// 接收上传的文件,需要使用
	router.POST("/upload", func9)

	router.Run(":8888")
}
