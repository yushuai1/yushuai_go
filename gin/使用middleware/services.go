package main

import (
	"github.com/gin-gonic/gin"
	"time"
	"log"
	"net/http"
)

func Logger() gin.HandlerFunc {
	return func(c *gin.Context) {
		t := time.Now()
		// 设置example变量到Context的Key中,通过Get等函数可以取得
		c.Set("example", "12345")
		// 发送request之前
		c.Next()
		// 发送request之后
		latency := time.Since(t)
		log.Print(latency)

		// 这个c.Write是ResponseWriter,我们可以获得状态等信息
		status := c.Writer.Status()
		log.Println(status)
	}
}

func main() {
	router := gin.Default()
	// 1
	router.Use(Logger())
	router.GET("/logger", func(c *gin.Context) {
		example := c.MustGet("example").(string)
		log.Println(example)
	})

	// 2
	// 下面测试BasicAuth()中间件登录认证
	//
	var secrets = gin.H{
		"foo":    gin.H{"email": "foo@bar.com", "phone": "123433"},
		"austin": gin.H{"email": "austin@example.com", "phone": "666"},
		"lena":   gin.H{"email": "lena@guapa.com", "phone": "523443"},
	}
	// Group using gin.BasicAuth() middleware
	// gin.Accounts is a shortcut for map[string]string
	authorized := router.Group("/admin", gin.BasicAuth(gin.Accounts{
		"foo":    "bar",
		"austin": "1234",
		"lena":   "hello2",
		"manu":   "4321",
	}))
	// 请求URL: 0.0.0.0:8888/admin/secrets
	authorized.GET("/secrets", func(c *gin.Context) {
		// get user, it was set by the BasicAuth middleware
		user := c.MustGet(gin.AuthUserKey).(string)
		log.Println(user,"-----------------")
		if secret, ok := secrets[user]; ok {
			c.JSON(http.StatusOK, gin.H{"user": user, "secret": secret})
		} else {
			c.JSON(http.StatusOK, gin.H{"user": user, "secret": "NO SECRET :("})
		}
	})

	router.Run(":8888")
}
