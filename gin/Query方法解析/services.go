package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"fmt"
	"time"
)

var TestInfo map[int]string = make(map[int]string)

var yu = 1

// 使用Query获取参数
func func6(c *gin.Context) {

	// 回复一个200 OK
	// 获取传入的参数
	name := c.Query("name")

	passwd := c.Query("passwd")

	c.String(http.StatusOK, "参数:%s %s  test6 OK", name, passwd)

}

func test(name string)  {

	yu++

	TestInfo[yu] = name

}
// 使用Query获取参数
func func7(c *gin.Context) {
	// 回复一个200 OK
	// 获取传入的参数
	name := c.PostForm("name")

	passwd := c.PostForm("passwd")

	test(name)

	if yu<0 {
		fmt.Println(passwd)
	}

	c.String(http.StatusOK, "参数:%s  test7 OK", name)
}

func main() {

	router := gin.Default()

	router.GET("/test6", func6)

	router.POST("/test7", func7)

	//gin.DisableConsoleColor()

	s := &http.Server{
		Addr:           ":8888",
		Handler:        router,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
	s.ListenAndServe()

	// 使用gin的Query参数形式,/test6?firstname=Jane&lastname=Doe

	//router.Run(":8888")
}
