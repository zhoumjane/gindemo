package v1

import (
	"fmt"
	"ginblog1/model"
	"ginblog1/utils/errmsg"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

// 添加文章
func AddArticle(c *gin.Context)  {
	var data model.Article
	_ = c.ShouldBindJSON(&data)
	code = model.CreateArticle(&data)
	c.JSON(http.StatusOK, gin.H{
		"status": code,
		"data": data,
		"message": errmsg.GetErrMsg(code),
	})
}

// 查询单个分类下的所有文章
func GetCategoryAllArticle(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	pageSize, _ := strconv.Atoi(c.Query("pagesize"))
	pageNum, _ := strconv.Atoi(c.Query("pagenum"))
	if pageSize == 0{
		pageSize = -1
	}
	if pageNum == 0{
		pageNum = -1
	}
	data, code := model.GetCategoryAllArticle(id, pageSize, pageNum)
	c.JSON(http.StatusOK, gin.H{
		"status": code,
		"data": data,
		"message": errmsg.GetErrMsg(code),
	})
}

// 查询文章详情
func GetArticle(c *gin.Context)  {
	id, _ := strconv.Atoi(c.Param("id"))
	data, code := model.GetArticle(id)
	c.JSON(http.StatusOK, gin.H{
		"status": code,
		"data": data,
		"message": errmsg.GetErrMsg(code),
	})
}
// 查询文章列表
func GetArticles(c *gin.Context)  {
	pageSize, _ := strconv.Atoi(c.Query("pagesize"))
	pageNum, _ := strconv.Atoi(c.Query("pagenum"))
	fmt.Println(pageSize)
	if pageSize == 0{
		pageSize = -1
	}
	if pageNum == 0{
		pageNum = 2
	}
	data := model.GetArticles(pageSize, pageNum)
	code = errmsg.SUCCESS
	c.JSON(http.StatusOK, gin.H{
		"status": code,
		"data": data,
		"message": errmsg.GetErrMsg(code),
	})
}
// 编辑文章
func EditArticle(c *gin.Context) {
	var data model.Article
	id, _ := strconv.Atoi(c.Param("id"))
	c.ShouldBindJSON(&data)
	code = model.EditArticle(id, &data)
	c.JSON(http.StatusOK, gin.H{
		"status": code,
		"message": errmsg.GetErrMsg(code),
	})
}
// 删除文章
func DeleteArticle(c *gin.Context)  {
	id, _ := strconv.Atoi(c.Param("id"))
	code = model.DeleteArticle(id)
	c.JSON(http.StatusOK, gin.H{
		"status": code,
		"message": errmsg.GetErrMsg(code),
	})
}