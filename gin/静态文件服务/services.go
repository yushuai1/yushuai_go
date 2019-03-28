package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main(){
	router := gin.Default()
	// 下面测试静态文件服务
	// 显示当前文件夹下的所有文件/或者指定文件
	router.StaticFS("/showDir", http.Dir("."))

	//获取当前文件的相对路径
	router.Static("/files", "./my")
	router.StaticFile("/image", "./cgo/yu.jpg")

	router.Run(":8888")
}