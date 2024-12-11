package main

import (
	"database/sql"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"log"
)

type student struct {
	Name    string `json:"name"`
	Id      string `json:"id"`
	Address string `json:"address"`
	Birth   string `json:"birth"`
	Gender  string `json:"gender"`
}

var db *sql.DB

func AddStudents(c *gin.Context) {
	var s student
	err := c.ShouldBindJSON(&s)
	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}
	_, err = db.Exec("INSERT INTO students (name,id,address,birth,gender) value (?,?,?,?,?)", s.Name, s.Id, s.Address, s.Birth, s.Gender)
	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
	} else {
		c.JSON(200, gin.H{
			"message": "Add successfully",
		})
	}
}
func ProfileStudents(c *gin.Context) {
	var s student
	err := c.ShouldBindJSON(&s)
	if err != nil {
		c.JSON(400, gin.H{
			"error": err,
		})
	} else {
		_, err = db.Exec("UPDATE students set name=?,id=?,address=?,birth=?,gender=? where id=?", s.Name, s.Id, s.Address, s.Birth, s.Gender, s.Id)
		if err != nil {
			c.JSON(400, gin.H{
				"error": err,
			})
		} else {
			c.JSON(200, gin.H{
				"message": "Profile successfully",
			})
		}
	}
}
func DeleteStudent(c *gin.Context) {
	id := c.Query("id")
	_, err := db.Exec("DELETE from students where id=?", id)
	if err != nil {
		c.JSON(400, gin.H{
			"error": err,
		})
	} else {
		c.JSON(200, gin.H{
			"message": "Delete successfully",
		})
	}
}
func SearchStudents(c *gin.Context) {
	var s student
	id := c.Query("id")
	row := db.QueryRow("SELECT * from students where id=?", id)
	err := row.Scan(&s.Name, &s.Id, &s.Address, &s.Birth, &s.Gender)
	if err != nil {
		c.JSON(400, gin.H{
			"error": err,
		})
	} else {
		c.JSON(200, s)
	}
}
func ClientDB() {
	dsn := "root:114514@tcp(127.0.0.1:3306)/mydb?charset=utf8mb4&parseTime=True&loc=Local"
	var err error
	db, err = sql.Open("mysql", dsn)
	if err != nil {
		log.Fatal(err)
	}

	err = db.Ping()
	if err != nil {
		log.Fatal(err)
	}
}
func main() {
	r := gin.Default()
	ClientDB()
	r.POST("/add", AddStudents)
	r.POST("/profile", ProfileStudents)
	r.GET("/delete", DeleteStudent)
	r.GET("/search", SearchStudents)
	r.Run(":8888")
}
