package model

import (
	"fmt"
	"ginblog1/dao"
	"ginblog1/utils/errmsg"
	"gorm.io/gorm"
)

var db *gorm.DB
var err error

func init()  {
	db, err = dao.InitDb()
	if err != nil{
		fmt.Println("数据库连接失败!")
		return
	}
	db.AutoMigrate(&User{}, &Article{}, &Category{}, &Student{}, &Book{})
}

type Article struct {
	gorm.Model
	Category Category `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	Title string `gorm:"type:varchar(100);not null" json:"title"`
	CategoryID int	`gorm:"type:int" json:"category_id"`
	Desc string `gorm:"type:varchar(200);not null" json:"desc"`
	Content string `gorm:"type:longtext;not null" json:"content"`
	Img string `gorm:"type:varchar(100);not null" json:"img"`
	User User `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	UserID int	`gorm:"type:int" json:"user_id"`
}

// 查询cid是否在category里存在
func CheckCid(id int) int {
	var category  Category
	db.Select("id").Where("id = ?", id).First(&category)
	fmt.Println(category.ID)
	if category.ID > 0{
		return errmsg.SUCCESS
	}
	return errmsg.ErrorCategoryNotExist
}

// 新增文章
func CreateArticle(data *Article) int {
	err := db.Create(&data).Error
	if err != nil{
		return errmsg.ERROR
	}
	return errmsg.SUCCESS
}

// 查询分类下的所有文章
func GetCategoryAllArticle(id int, pageSize int, pageNum int) ([]Article, int) {
	var cateArtList []Article
	err := db.Preload("Category").Limit(pageSize).Offset((pageNum - 1)* pageSize).Where("CategoryID = ?", id).Find(&cateArtList).Error
	if err != nil {
		return nil, errmsg.ErrorCategoryNotExist
	}
	return cateArtList, errmsg.SUCCESS
}

// 查询文章详情
func GetArticle(id int) (Article, int) {
	var art Article
	err := db.Preload("Category").Where("id = ?", id).First(&art).Error
	if err != nil{
		return art, errmsg.ErrorArtNotExist
	}
	return art, errmsg.SUCCESS
}
// 查询文章列表
func GetArticles(pageSize int, pageNum int) []Article {
	var art  []Article
	err = db.Preload("Category").Preload("User").Limit(pageSize).Offset((pageNum - 1) * pageSize).Find(&art).Error
	//err = db.Preload("Category").Limit(pageSize).Offset((pageNum - 1) * pageSize).Find(&art).Error
	if err != nil && err != gorm.ErrRecordNotFound{
		return nil
	}
	return art
}

// 编辑文章
func EditArticle(id int, data *Article) int {
	var art Article
	var maps = make(map[string]interface{})
	maps["title"] = data.Title
	maps["cid"] = data.CategoryID
	maps["desc"] = data.Desc
	maps["content"] = data.Content
	maps["img"] = data.Img
	err = db.Model(&art).Where("id = ?", id).Updates(maps).Error
	if err != nil{
		return errmsg.ERROR
	}
	return errmsg.SUCCESS
}

// 删除文章
func DeleteArticle(id int) int {
	var art Article
	err = db.Where("id = ?", id).Delete(&art).Error
	if err != nil{
		return errmsg.ERROR
	}
	return errmsg.SUCCESS
}