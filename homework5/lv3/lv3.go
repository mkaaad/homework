package main

import (
	"bufio"
	"encoding/json"
	"os"

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
		err:=AddData(s)
		if err!=nil{
		c.JSON(400,gin.H{
			"error": err.Error(),
		})
		}else{
		c.JSON(200, gin.H{
			"message":     "Add successfully",
			"studentINFO":	students[s.Id],
		})
	}
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
		err= UpdateData()
		if err!=nil{
			c.JSON(400,gin.H{
			"error": err.Error(),
			})
		}else{	
			c.JSON(200, gin.H{
			"message":     "Profile successfully",
			"studentINFO":	students[s.Id],
		})
	}
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
func ReadFile(){
	file,err:=os.OpenFile("data.json",os.O_CREATE|os.O_RDONLY,0666)
	if err!=nil{
		panic(err)
	}
	defer file.Close()

	scanner:=bufio.NewScanner(file)
	for scanner.Scan(){
		var s student
		err=json.Unmarshal(scanner.Bytes(),&s)
		if err!=nil{
			panic(err)
		}
		students[s.Id]=&s
	}
}
func AddData(s student)error{
	file,err:=os.OpenFile("data.json",os.O_APPEND|os.O_WRONLY,0666)
	if err!=nil{
		panic(err)
	}
	defer file.Close()
	
	jsonData,err:=json.Marshal(s)
	if err!=nil{
		return err
	}
	jsonData=append(jsonData, '\n')

	_,err=file.Write(jsonData)
	if err!=nil{
		return err
	}
	return nil
}
func UpdateData()error{
	file,err:=os.OpenFile("data.json",os.O_WRONLY|os.O_TRUNC,0666)
	if err!=nil{
		panic(err)
	}
	defer file.Close()

	for _,s:= range students{
		err:=AddData(*s)
		if err!=nil{
			return err
		}
	}
	return nil
}
func main() {
	r := gin.Default()
	students = make(map[string]*student)
	ReadFile()
	r.POST("/add", AddStudents)
	r.POST("/profile", ProfileStudents)
	r.GET("/search", SearchStudents)
	r.Run(":8888")
}
