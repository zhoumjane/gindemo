package v1

import (
	"ginblog1/middleware"
	"ginblog1/model"
	"ginblog1/utils/errmsg"
	"github.com/gin-gonic/gin"
	"net/http"
)
func Login(c *gin.Context)  {
	var data model.User
	var code int
	var token string
	c.ShouldBindJSON(&data)
	code = model.CheckLogin(data.Username, data.Password)
	if code == errmsg.SUCCESS{
		token, code = middleware.SetToken(data.Username, data.Password)
	}
	c.JSON(http.StatusOK, gin.H{
		"status": code,
		"message": errmsg.GetErrMsg(code),
		"token": token,
	})
}