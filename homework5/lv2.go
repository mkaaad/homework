package main

import (
	"github.com/gin-gonic/gin"
)

type student struct {
	Name    string `json:"name"`
	Id      string `json:"id"`
	Address string `json:"address"`
	Birth   string `json:"birth"`
	Gender  string `json:"gender"`
}

var students map[string]*student

func AddStudents(c *gin.Context) {
	var s student
	err := c.ShouldBindJSON(&s)
	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
	} else if students[s.Id]!=nil{
		c.JSON(400,gin.H{
			"error":	"Already exist",
		})
	} else {
		students[s.Id] = &s
		c.JSON(200, gin.H{
			"message":     "Add successfully",
			"studentINFO":	students[s.Id],
		})
	}
}
func ProfileStudents(c *gin.Context) {
	var s student
	err := c.ShouldBindJSON(&s)
	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
	} else if students[s.Id] != nil {
		students[s.Id] = &s
		c.JSON(200, gin.H{
			"message":     "Profile successfully",
			"studentINFO":	students[s.Id],
		})
	} else {
		c.JSON(400, gin.H{
			"error": "Can not find student",
		})
	}
}
func SearchStudents(c *gin.Context) {
	id := c.Query("id")
	if students[id] != nil {
		c.JSON(200, students[id])
	}

}
func main() {
	r := gin.Default()
	students = make(map[string]*student)
	r.POST("/add", AddStudents)
	r.POST("/profile", ProfileStudents)
	r.GET("/search", SearchStudents)
	r.Run(":8888")
}
