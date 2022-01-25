package middleware

import (
	"ginblog1/utils"
	"ginblog1/utils/errmsg"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
	"time"
)

var JwtKey  = []byte(utils.JwtKey)
var code int

type MyClaims struct {
	Username string `json:"username"`
	Password string `json:"password"`
	jwt.StandardClaims
}

// 生成Token
func SetToken(username string, password string) (string, int) {
	expireTime := time.Now().Add(24 * time.Hour)
	SetClaims := MyClaims{
		Username: username,
		Password: password,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expireTime.Unix(),
			Issuer:    "ginblog",
		},
	}
	reqClaim := jwt.NewWithClaims(jwt.SigningMethodHS256, SetClaims)
	token, err := reqClaim.SignedString(JwtKey)
	if err != nil{
		return "", errmsg.ERROR
	}
	return token, errmsg.SUCCESS
}
// 验证Token
func CheckToken(token string) (*MyClaims, int) {
	setToken, _ := jwt.ParseWithClaims(token, &MyClaims{}, func(token *jwt.Token) (i interface{}, err error) {
		return JwtKey, nil
	})
	if key, ok := setToken.Claims.(*MyClaims); ok && setToken.Valid{
		return key, errmsg.SUCCESS
	}else {
		return nil, errmsg.ERROR
	}
}
// jwt中间件
func JwtToken() gin.HandlerFunc {
	return func(context *gin.Context) {
		tokenHeader := context.Request.Header.Get("Authorization")
		if tokenHeader == ""{
			code = errmsg.ErrorTokenNotExist
			context.JSON(http.StatusOK, gin.H{
				"code": code,
				"message": errmsg.GetErrMsg(code),
			})
			context.Abort()
			return
		}
		code = errmsg.SUCCESS
		checkToken := strings.SplitN(tokenHeader, " ", 2)
		if len(checkToken) != 2 && checkToken[0] != "Bearer" {
			code = errmsg.ErrorTokenTypeWrong
			context.JSON(http.StatusOK, gin.H{
				"code": code,
				"message": errmsg.GetErrMsg(code),
			})
			context.Abort()
			return
		}
		key, tCode := CheckToken(checkToken[1])
		if tCode == errmsg.ERROR{
			code = errmsg.ErrorTokenWrong
			context.JSON(http.StatusOK, gin.H{
				"code": code,
				"message": errmsg.GetErrMsg(code),
			})
			context.Abort()
			return
		}
		if time.Now().Unix() > key.ExpiresAt {
			code = errmsg.ErrorTokenRuntime
			context.JSON(http.StatusOK, gin.H{
				"code": code,
				"message": errmsg.GetErrMsg(code),
			})
			context.Abort()
			return
		}
		context.Set("username", key.Username)
		context.Next()
	}
}