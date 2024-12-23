package handlers

import (
	"homework7/pkg/database"
	"homework7/pkg/models"
	"github.com/gin-gonic/gin"
)
func GetMessage(c *gin.Context){
	var m models.Messages
	var message []models.Messages
	if !CheckLogin(c){
		return
	}
	id := c.Query("id")
	rows,err:= database.Db.Query("SELECT * from messages where (id=? Or parent_id=?) And is_deleted!=1", id,id)        
	if err != nil {
                c.JSON(500, gin.H{
                        "message": err,
                })
		return
        }
	defer rows.Close()
	for rows.Next(){
        err = rows.Scan(&m.Id,&m.UserId,&m.Context,&m.Created_at,&m.Updated_at,&m.IsDeleted,&m.ParentId)
	if err!=nil{
		c.JSON(500, gin.H{
                        "message": err,
                })

			return
		}
		message = append(message, m)
	}
	c.JSON(200,message)
}
func PostMessage(c *gin.Context){
	var m models.Messages
	m.ParentId=-1
	if !CheckLogin(c){
		return
	}
	err:=c.ShouldBindJSON(&m)
	if err != nil {
                c.JSON(400, gin.H{
                        "message": err,
                })
                return
        }
	if m.ParentId!=-1{
	_, err = database.Db.Exec("INSERT INTO messages (user_id,context,parent_id) value (?,?,?)",  LoginUser.Id, m.Context,m.ParentId)
	}else{
	_,err=database.Db.Exec("INSERT INTO messages (user_id,context) value (?,?)",  m.UserId, m.Context)
}
        if err != nil {
                c.JSON(500, gin.H{
                        "message": err.Error(),
                })
        } else {
                c.JSON(200, gin.H{
                        "message": "Post successfully",
                })
        }
}
func DeleteMessage(c *gin.Context){
	if !CheckLogin(c){
		return
	}
	id := c.Query("id")
	_, err := database.Db.Exec("UPDATE messages set is_deleted=1 where (id=? Or parent_id=?) And user_id=?",id,id,LoginUser.Id)
	if err != nil {
                c.JSON(500, gin.H{
                        "message": err,
                })
        } else {
                c.JSON(200, gin.H{
                        "message": "Delete successfully",
                })
        }

}
func LikeMessage(c *gin.Context){
	if !CheckLogin(c){
		return
	}
	messageId := c.Query("id")
	rows, err := database.Db.Query("SELECT * from like where user_id=? And message_id=?",  LoginUser.Id,messageId)
	if err!=nil{
		c.JSON(500,gin.H{
			"message":err,
		})
		return
	}
	if rows.Next(){
		_,err=database.Db.Exec("Delete from like where")
		if err!=nil{
			c.JSON(500,gin.H{
				"message":err,
			})
			return
		}else{
			c.JSON(200,gin.H{
				"message":"Cancel like successfully",
		})
			return
		}
	}


	_,err=database.Db.Exec("INSERT INTO like (user_id,message_id) value (?,?)",LoginUser.Id,messageId)
	if err!=nil{
		c.JSON(500,gin.H{
			"message":err,
		})
		return
	}
	c.JSON(200,gin.H{
		"message":"Like successfully",
	})
}
