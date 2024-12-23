package main

import (
	"homework7/pkg/database"
	"homework7/pkg/handlers"

	"github.com/gin-gonic/gin"
)
func main(){
	r:=gin.Default()
	database.ClientDB()
	r.POST("/login",handlers.Login)
	r.POST("/register",handlers.Register)
	r.POST("/profile",handlers.ProfileUser)
	r.GET("/getmessage",handlers.GetMessage)
	r.POST("/postmessage",handlers.PostMessage)
	r.GET("/deletemessage",handlers.DeleteMessage)
	r.GET("/likemessage",handlers.LikeMessage)
	r.Run(":8888")
}
