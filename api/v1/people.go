package v1

import (
	"fmt"
	"ginblog1/model"
	"ginblog1/utils/errmsg"
	"ginblog1/utils/validator"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type BookJson struct {
	Name string `json:"name"`
	Students []int `json:"students"`
}

// 创建学生
func AddStudent(c *gin.Context)  {
	var  data  model.Student
	var msg string
	_ = c.ShouldBindJSON(&data)
	msg, code = validator.Validate(data)
	if code != errmsg.SUCCESS {
		c.JSON(http.StatusOK, gin.H{
			"status":  code,
			"message": msg,
		})
		return
	}
	if code == errmsg.SUCCESS{
		code = model.CreateStudent(&data)
	}
	c.JSON(http.StatusOK, gin.H{
		"status": code,
		"data": data,
		"message": errmsg.GetErrMsg(code),
	})
}

// 创建书籍
func AddBook(c *gin.Context)  {
	var  data  model.Book
	var jsonData BookJson
	var msg string
	_ = c.ShouldBindJSON(&jsonData)
	studentsId := jsonData.Students
	code = model.CheckStudentIsExist(studentsId)
	if code == errmsg.ErrorStudentNameNotExist{
		c.JSON(http.StatusOK, gin.H{
			"status": code,
			"data": data,
			"message": errmsg.GetErrMsg(code),
		})
		return
	}
	data.Name = jsonData.Name
	msg, code = validator.Validate(data)
	if code != errmsg.SUCCESS {
		c.JSON(http.StatusOK, gin.H{
			"status":  code,
			"message": msg,
		})
		return
	}
	if code == errmsg.SUCCESS{
		code = model.CreateBook(&data)
	}
	c.JSON(http.StatusOK, gin.H{
		"status": code,
		"data": data,
		"message": errmsg.GetErrMsg(code),
	})
	model.CreateForeignKeyStudent(jsonData.Students, int(data.ID))
}

// 查询一个学生拥有多少书籍
func GetStudentBooks(c *gin.Context)  {
	user, _ := c.Get("username")
	fmt.Println(user)
	id, _ := strconv.Atoi(c.Param("id"))
	data, code := model.GetStudentBooks(id)
	c.JSON(http.StatusOK, gin.H{
		"status": code,
		"data": data,
		"message": errmsg.GetErrMsg(code),
	})
}
