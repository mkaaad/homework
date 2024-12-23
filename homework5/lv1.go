package main

import "github.com/gin-gonic/gin"

func main() {
	r := gin.Default()
	r.GET("/ping", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{
			"message": "pong",
		})
	})
	r.GET("/echo", func(ctx *gin.Context) {
		message := ctx.Query("message")
		ctx.JSON(200, gin.H{
			"message": message,
		})
	})

	r.Run(":8888")
}
