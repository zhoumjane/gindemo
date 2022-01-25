package model

import (
	"encoding/base64"
	//"fmt"
	"ginblog1/utils/errmsg"
	"golang.org/x/crypto/scrypt"
	//"github.com/go-playground/validator"
	"gorm.io/gorm"
	"log"
	//"reflect"
	//"strings"

	//"reflect"
	//"strings"
)


type User struct {
	gorm.Model
	Username string `gorm:"unique;type:varchar(20);not null" json:"username" validate:"required,min=6,max=20" form:"username"`
	Password string `gorm:"type:varchar(20);not null" json:"password" validate:"required,min=6,max=20" form:"password"`
	Role int 	`gorm:"type:int" json:"role" validate:"required,gte=0" form:"role"`
}

// 查询用户是否存在
func CheckUser(username string) (code int) {
	var users  User
	db.Select("id").Where("username = ?", username).First(&users)
	if users.ID > 0{
		return errmsg.ErrorUsernameUsed
	}
	return errmsg.SUCCESS
}


// 新增用户
func CreateUser(data *User) int {
	   err = db.Create(&data).Error
	   if err != nil{
		   return errmsg.ErrorUsernameUsed
	   }
	   return errmsg.SUCCESS
}

// 查询用户列表
func GetUsers(pageSize int, pageNum int) []User  {
	var users  []User
	err = db.Limit(pageSize).Offset((pageNum-1) * pageSize).Find(&users).Error
	if err != nil && err != gorm.ErrRecordNotFound{
		return nil
	}
	return users
}

// 编辑用户
func EditUser(id int, data *User) int {
	var user User
	var maps = make(map[string]interface{})
	maps["username"] = data.Username
	maps["role"] = data.Role
	err = db.Model(&user).Where("id = ?", id).Updates(maps).Error
	if err != nil{
		return errmsg.ERROR
	}
	return errmsg.SUCCESS
}

// 删除用户
func DeleteUser(id int) int {
	var user User
	err = db.Where("id = ?", id).Delete(&user).Error
	if err != nil{
		return errmsg.ERROR
	}
	return errmsg.SUCCESS
}
// 密码加密
func (u *User) BeforeSave(tx *gorm.DB) (err error){
	u.Password = ScryptPw(u.Password)
	return nil
}
func ScryptPw(password string) string {
	const KeyLen  = 10
	salt := make([]byte, 8)
	salt = []byte{12, 32, 4, 6, 66, 22, 222, 11}
	hashPw, err := scrypt.Key([]byte(password), salt, 16384,8, 1, KeyLen)
	if err != nil{
		log.Fatal(err)
	}
	Fpw := base64.StdEncoding.EncodeToString(hashPw)
	return Fpw
}

// 登录验证
func CheckLogin(username string, password string) int {
	var user User
	//db.Where("username = ?", username).First(&user)
	db.Where("username = ?", username).First(&user)
	if user.ID == 0{
		return errmsg.ErrorUserNotExist
	}
	if ScryptPw(password) != user.Password {
		return errmsg.ErrorPasswordWrong
	}
	if user.Role != 1{
		return errmsg.ErrorUserNoRight
	}
	return errmsg.SUCCESS
}