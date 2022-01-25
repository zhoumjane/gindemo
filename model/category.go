package model

import (
	"fmt"
	"ginblog1/utils/errmsg"
	"gorm.io/gorm"
)

type Category struct {
	ID   uint   `gorm:"primary_key;auto_increment" json:"id"`
	Name string `gorm:"index:idx_name;unique;not null;type:varchar(20)" json:"name"`
}

// 查询分类是否存在
func CheckCategory(name string) (code int) {
	var cate Category
	db.Select("id").Where("name = ?", name).First(&cate)
	fmt.Println(cate.ID)
	if cate.ID > 0{
		return errmsg.ErrorCategoryUsed
	}
	return errmsg.SUCCESS
}

// 新增分类
func CreateCategory(data *Category) int {
	err := db.Create(&data).Error
	if err != nil{
		return errmsg.ERROR
	}
	return errmsg.SUCCESS
}

// 查询分类信息
func GetCateGory(pageSize int, pageNum int) []Category {
	var cate []Category
	err = db.Limit(pageSize).Offset((pageNum - 1) * pageSize).Find(&cate).Error
	if err != nil && err != gorm.ErrRecordNotFound{
		return nil
	}
	return cate
}

// 编辑分类信息
func EditCategory(id int, date *Category) int {
	var cate Category
	var maps = make(map[string]interface{})
	maps["name"] = date.Name
	err = db.Model(&cate).Where("id = ?", id).Updates(maps).Error
	if err != nil{
		return errmsg.ERROR
	}
	return errmsg.SUCCESS
}

// 删除分类
func DeleteCategory(id int) int {
	var cate Category
	err = db.Where("id = ?", id).Delete(&cate).Error
	if err != nil{
		return errmsg.ERROR
	}
	return errmsg.SUCCESS
}