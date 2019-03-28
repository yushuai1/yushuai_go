package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// router GROUP - GET测试
func func10(c *gin.Context) {
	c.String(http.StatusOK, "test10 OK")
}
func func11(c *gin.Context) {
	c.String(http.StatusOK, "test11 OK")
}

// router GROUP - POST测试
func func12(c *gin.Context) {
	c.String(http.StatusOK, "test12 OK")
}
func func13(c *gin.Context) {
	c.String(http.StatusOK, "test13 OK")
}

func main() {
	router := gin.Default()
	// router Group是为了将一些前缀相同的URL请求放在一起管理
	group1 := router.Group("/g1")
	group1.GET("/read1", func10)
	group1.GET("/read2", func11)

	group2 := router.Group("/g2")
	group2.POST("/write1", func12)
	group2.POST("/write2", func13)

	router.Run(":8888")
}
