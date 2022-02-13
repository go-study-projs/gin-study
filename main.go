package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	// 分组 v1
	v1 := router.Group("/v1")
	{
		v1.POST("/login", loginEndpoint)
		v1.POST("/submit", submitEndpoint)
		v1.POST("/send", sendEndpoint)
		v1.POST("/aaa", aaaaEndpoint)
		v1.POST("/bbb", bbbbEndpoint)
		v1.POST("/ccc", ccccEndpoint)
		v1.POST("/ddd", ddddEndpoint)
	}

	// 分组 v2
	v2 := router.Group("/v2")
	{
		v2.POST("/login", loginEndpoint2)
		v2.POST("/submit", submitEndpoint2)
		v2.POST("/send", sendEndpoint2)
		v2.POST("/aaa", aaaaEndpoint2)
		v2.POST("/bbb", bbbbEndpoint2)
		v2.POST("/ccc", ccccEndpoint2)
		v2.POST("/ddd", ddddEndpoint2)

	}
	v2Path := v2.Group("/path")
	v2Path.POST("/aaa", ddddEndpoint3)

	router.Run(":8080")
}

func ddddEndpoint3(context *gin.Context) {

}

func ddddEndpoint2(context *gin.Context) {

}

func ccccEndpoint2(context *gin.Context) {

}

func bbbbEndpoint2(context *gin.Context) {

}

func aaaaEndpoint2(context *gin.Context) {

}

func sendEndpoint2(context *gin.Context) {

}

func submitEndpoint2(context *gin.Context) {

}

func loginEndpoint2(context *gin.Context) {

}

func ddddEndpoint(context *gin.Context) {

}

func ccccEndpoint(context *gin.Context) {

}

func bbbbEndpoint(context *gin.Context) {

}

func aaaaEndpoint(context *gin.Context) {

}

func sendEndpoint(context *gin.Context) {

}

func submitEndpoint(context *gin.Context) {

}

func loginEndpoint(context *gin.Context) {

}
