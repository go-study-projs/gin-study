package main

import (
	"github.com/gin-gonic/gin"
	"io"
	"os"
)

func main() {
	// 不启动颜色控制
	gin.DisableConsoleColor()

	// 创建日志文件
	f, _ := os.Create("gin.log")
	gin.DefaultWriter = io.MultiWriter(f)

	// 使用如下代码，可以实现日志文件记录和 console 同步显示
	// gin.DefaultWriter = io.MultiWriter(f, os.Stdout)

	router := gin.Default()
	router.GET("/ping", func(c *gin.Context) {
		c.String(200, "pong")
	})

	router.Run(":8080")
}
