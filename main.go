package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	router := gin.Default()
	// 定义一个Get类型的服务接口
	// 可以匹配 /user/john，但是不能匹配 /user/ 和 /user 请求
	router.GET("/user/:name", func(c *gin.Context) {
		name := c.Param("name")
		c.String(http.StatusOK, "Hello %s", name)
	})

	// 可以匹配 /user/john/ 和 /user/john/send
	router.GET("/user/:name/*action", func(c *gin.Context) {
		name := c.Param("name")
		action := c.Param("action")
		message := name + " is " + action
		c.String(http.StatusOK, message)
	})

	// 对于每个匹配的请求，上下文将保存路由定义
	router.POST("/user/:name/*action", func(c *gin.Context) {
		b := c.FullPath() == "/user/:name/*action" // true
		c.String(http.StatusOK, "b is true? : %v", b)
	})

	router.Run(":8080")
}
