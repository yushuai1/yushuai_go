package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"fmt"
)

// func3: 处理带参数的path-GET
func func3(c *gin.Context)  {
	// 回复一个200 OK
	// 获取传入的参数
	yu := c.Param("yu")
	shuai := c.Param("shuai")
	fmt.Printf("参数:%s %s", yu,shuai)
	c.String(http.StatusOK, "参数:%s test3 OK", yu)
}
// func4: 处理带参数的path-POST
func func4(c *gin.Context)  {
	// 回复一个200 OK
	// 获取传入的参数
	name := c.Param("name")
	passwd := c.Param("passwd")
	c.String(http.StatusOK, "参数:%s %s  test4 OK", name, passwd)
}
// func5: 注意':'和'*'的区别
func func5(c *gin.Context)  {
	// 回复一个200 OK
	// 获取传入的参数
	name := c.Param("name")
	passwd := c.Param("passwd")
	c.String(http.StatusOK, "参数:%s %s  test5 OK", name, passwd)
}

func main(){
	router := gin.Default()
	// TODO:注意':'必须要匹配,'*'选择匹配,即存在就匹配,否则可以不考虑
	router.GET("/test3/:yu/:shuai", func3)
	router.POST("/test4/:name/:passwd", func4)
	router.GET("/test5/:name/*passwd", func5)

	router.Run(":8888")
}
