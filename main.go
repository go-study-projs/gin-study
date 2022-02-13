package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// Login Binding from JSON
type Login struct {
	User     string `form:"user" json:"user" xml:"user"  binding:"required"`
	Password string `form:"password" json:"password" xml:"password" binding:"required"`
}

func main() {
	router := gin.Default()

	// 绑定 JSON 类型参数 ({"user": "manu", "password": "123"})
	router.POST("/loginJSON", func(c *gin.Context) {
		var json Login
		if err := c.ShouldBindJSON(&json); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		if json.User != "manu" || json.Password != "123" {
			c.JSON(http.StatusUnauthorized, gin.H{"status": "unauthorized"})
			return
		}

		c.JSON(http.StatusOK, gin.H{"status": "you are logged in"})
	})

	// Example for binding XML (
	//  <?xml version="1.0" encoding="UTF-8"?>
	//  <root>
	//    <user>user</user>
	//    <password>123</password>
	//  </root>)
	router.POST("/loginXML", func(c *gin.Context) {
		var xml Login
		if err := c.ShouldBindXML(&xml); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		if xml.User != "manu" || xml.Password != "123" {
			c.JSON(http.StatusUnauthorized, gin.H{"status": "unauthorized"})
			return
		}

		c.JSON(http.StatusOK, gin.H{"status": "you are logged in"})
	})

	// 绑定 HTML 形式的参数 (user=manu&password=123)
	router.POST("/loginForm", func(c *gin.Context) {
		var form Login
		// This will infer what binder to use depending on the content-type header.
		if err := c.ShouldBind(&form); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		if form.User != "manu" || form.Password != "123" {
			c.JSON(http.StatusUnauthorized, gin.H{"status": "unauthorized"})
			return
		}

		c.JSON(http.StatusOK, gin.H{"status": "you are logged in"})
	})

	// 启动服务并监听 8080 端口
	router.Run(":8080")
}

//此外，在进行参数模型绑定时，Gin 还提供了两组绑定方法 MustBindWith 和 ShouldBindWith，
//一种是必须绑定模式，一种是应该绑定模式。
//前者在发生解析错误时，会直接触发错误返回，后者需要开发者自己进行错误处理。
//相关的方法如下所示：必须绑定模式有 Bind,  BindJSON,  BindXML,  BindQuery,  BindYAML，
//应该绑定模式有 ShouldBind,  ShouldBindJSON,  ShouldBindXML,  ShouldBindQuery,  ShouldBindYAML。
