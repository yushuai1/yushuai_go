package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
	"log"
)


func main(){
	router:=gin.Default()

	rg := router.Group("/yu",Validate())
	//router.Use(Validate())  //使用validate()中间件身份验证
	rg.GET("/shuai",Service1)
	rg.GET("/jia",Service2)

	rg1 := router.Group("/ff",Logger(),Validate())
	rg1.GET("/shaui",Service3)
	router.Run(":8989")  //localhost:8989/
}

func Service1(c *gin.Context){
	c.JSON(http.StatusOK,gin.H{"message":"你好，欢迎你1"})
}
func Service2(c *gin.Context){
	c.JSON(http.StatusOK,gin.H{"message":"你好，欢迎你2"})
}
func Service3(c *gin.Context){
	c.JSON(http.StatusOK,gin.H{"message":"你好，欢迎你3"})
}

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



func Validate() gin.HandlerFunc{
	return func(c *gin.Context){
		//这一部分可以替换成从session/cookie中获取，
		username:=c.Query("username")
		password:=c.Query("password")

		if username=="ft" && password =="123"{
			c.JSON(http.StatusOK,gin.H{"message":"身份验证成功"})
			c.Next()  //该句可以省略，写出来只是表明可以进行验证下一步中间件，不写，也是内置会继续访问下一个中间件的
		}else {
			c.Abort()
			c.JSON(http.StatusUnauthorized,gin.H{"message":"身份验证失败"})
			return// return也是可以省略的，执行了abort操作，会内置在中间件defer前，return，写出来也只是解答为什么Abort()之后，还能执行返回JSON数据
		}
	}
}