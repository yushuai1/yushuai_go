package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	router := gin.Default()
	// 下面测试加载HTML: LoadHTMLTemplates
	// 加载templates文件夹下所有的文件
	router.LoadHTMLGlob("templates/*")
	// 或者使用这种方法加载也是OK的: router.LoadHTMLFiles("templates/template1.html", "templates/template2.html")
	router.GET("/index", func(c *gin.Context) {
		// 注意下面将gin.H参数传入index.tmpl中!也就是使用的是index.tmpl模板
		c.HTML(http.StatusOK, "index.tmpl", gin.H{
			"title": "GIN: 测试加载HTML模板",
		})
	})

	router.Run(":8888")
}
