package handlers

import (
	"homework7/pkg/models"
	"homework7/pkg/database"
	"github.com/gin-gonic/gin"
)
var isLogin bool
func CheckLogin(c *gin.Context) bool {
        if !isLogin {
                c.JSON(401, gin.H{
                        "message": "please login first",
                })
                return false
        } else {
                return true
        }
}
func Login(c *gin.Context) {
        if CheckLogin(c) {
                c.JSON(200, gin.H{
                        "message": "Login successfully",
                })
                return
        }
        var u models.User 
        err := c.ShouldBindJSON(&u)
        if err != nil {
                c.JSON(400, gin.H{
                        "message": err,
                })
                return
        }
	if u.UserName=="" || u.Password==""{
		c.JSON(400,gin.H{
			"message":"Id and password can not be void",
		})
		return
	}
        rows, err := database.Db.Query("SELECT * from users where id=? and password=?", u.Id, u.Password)
        if !rows.Next() {
                c.JSON(500, gin.H{
                        "message": "user does not exist or passwd is wrong",
                })
                return
        }
        if err != nil {
                c.JSON(400, gin.H{
                        "message": err,
                })
                return
        }
        err = rows.Scan(&u.Id, &u.NickName,&u.Password, &u.Created_at,&u.Updated_at)
        if err != nil {
                c.JSON(400, gin.H{
                        "message": err,
                })
        } else {
                isLogin = true
                c.JSON(200,u)
        }
}
func Register(c *gin.Context) {
        var u models.User
        err := c.ShouldBindJSON(&u)
        if err != nil {
                c.JSON(400, gin.H{
                        "message": err,
                })
                return
        }
	if  u.Password==""||u.UserName==""{
		c.JSON(400,gin.H{
			"message":"Username and password can not be void",
		})
		return
	}
        _, err = database.Db.Exec("INSERT INTO users (id,nickname,password,created_at,updated_at) value (,?,?,?,?)", u.NickName,u.Password, u.Created_at,u.Updated_at)
        if err != nil {
                c.JSON(400, gin.H{
                        "message": err,
                })
        } else {
                c.JSON(200, gin.H{
                        "message": "register successfully",
                })
        }
}

