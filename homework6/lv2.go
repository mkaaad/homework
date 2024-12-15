package main

import (
	"database/sql"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"log"
)

var isLogin bool
var db *sql.DB

type student struct {
	Name    string `json:"name"`
	Id      string `json:"id"`
	Address string `json:"address"`
	Birth   string `json:"birth"`
	Gender  string `json:"gender"`
}
type user struct {
	Username   string `json:"username"`
	Passwd     string `json:"passwd"`
	Permission string `json:"permission"`
}

func CheckLogin(c *gin.Context) bool {
	if !isLogin {
		c.JSON(400, gin.H{
			"error": "please login first",
		})
		return false
	} else {
		return true
	}
}

func Login(c *gin.Context) {
	if CheckLogin(c) {
		c.JSON(400, gin.H{
			"error": "already login",
		})
		return
	}
	var u user
	err := c.ShouldBindJSON(&u)
	if err != nil {
		c.JSON(400, gin.H{
			"error": err,
		})
		return
	}
	rows, err := db.Query("SELECT * from user where username=? and passwd=?", u.Username, u.Passwd)
	if !rows.Next() {
		c.JSON(400, gin.H{
			"wrong": "user does not exist or passwd is wrong",
		})
		return
	}
	if err != nil {
		c.JSON(400, gin.H{
			"error": err,
		})
		return
	}
	err = rows.Scan(&u.Username, &u.Passwd, &u.Permission)
	if err != nil {
		c.JSON(400, gin.H{
			"error": err,
		})
	} else {
		isLogin = true
		c.JSON(200, gin.H{
			"message": "login successfully",
		})
	}
}
func Register(c *gin.Context) {
	var u user
	err := c.ShouldBindJSON(&u)
	if err != nil {
		c.JSON(400, gin.H{
			"error": err,
		})
		return
	}
	_, err = db.Exec("INSERT INTO user (username,passwd,permission) value (?,?,?)", u.Username, u.Passwd, u.Permission)
	if err != nil {
		c.JSON(400, gin.H{
			"error": err,
		})
	} else {
		c.JSON(200, gin.H{
			"message": "register successfully",
		})
	}
}
func AddStudents(c *gin.Context) {
	var s student
	if !CheckLogin(c) {
		return
	}
	err := c.ShouldBindJSON(&s)
	if err != nil {
		c.JSON(400, gin.H{
			"error": err,
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
	if !CheckLogin(c) {
		return
	}

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
	if !CheckLogin(c) {
		return
	}

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
	if !CheckLogin(c) {
		return
	}
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
	isLogin = false
	r.POST("/login", Login)
	r.POST("/register", Register)
	r.POST("/add", AddStudents)
	r.POST("/profile", ProfileStudents)
	r.GET("/delete", DeleteStudent)
	r.GET("/search", SearchStudents)
	r.Run(":8888")
}
