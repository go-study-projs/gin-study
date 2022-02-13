package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	//默认使用了日志中间件和 Recovery 中间件
	//	r := gin.Default()

	// 创建空白路由
	r := gin.New()

	// 集成日志中间件
	// 默认 gin.DefaultWriter = os.Stdout
	r.Use(gin.Logger())

	// 集成 Recovery 中间件
	r.Use(gin.Recovery())

	// Per route middleware, you can add as many as you desire.
	r.GET("/benchmark", MyBenchLogger(), benchEndpoint)

	// 认证分组
	// authorized := r.Group("/", AuthRequired())
	// 上面的书写方法等价于下面的编码方式
	authorized := r.Group("/")
	authorized.Use(AuthRequired())
	{
		authorized.POST("/login", loginEndpoint)
		authorized.POST("/submit", submitEndpoint)
		authorized.POST("/read", readEndpoint)
	}

	// 启动服务监听 8080 端口
	r.Run(":8080")
}

func readEndpoint(context *gin.Context) {

}

func submitEndpoint(context *gin.Context) {

}

func loginEndpoint(context *gin.Context) {

}

func AuthRequired() gin.HandlerFunc {

	return nil
}

func benchEndpoint(context *gin.Context) {

}

func MyBenchLogger() gin.HandlerFunc {
	return nil
}
