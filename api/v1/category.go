package v1

import (
	"ginblog1/model"
	"ginblog1/utils/errmsg"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

// 添加分类
func AddCategory(c *gin.Context) {
	var data model.Category
	_ = c.ShouldBindJSON(&data)
	//code = model.CheckCategory(data.Name)
	//if code == errmsg.SUCCESS{
	//	model.CreateCategory(&data)
	//}
	//if code == errmsg.ErrorCategoryUsed{
	//	code = errmsg.ErrorCategoryUsed
	//}
	code = model.CreateCategory(&data)
	c.JSON(http.StatusOK, gin.H{
		"status": code,
		"data": data,
		"message": errmsg.GetErrMsg(code),
	})

}

// 查询单个分类下的所有文章
//func GetCategoryAllArticle(c *gin.Context) int {
//	return 0
//}

// 查询分类列表
func GetCategorys(c *gin.Context)  {
		pageSize, _ := strconv.Atoi(c.Query("pagesize"))
		pageNum, _ := strconv.Atoi(c.Query("pagenum"))
		if pageSize == 0{
			pageSize = -1
		}
		if pageNum == 0{
			pageNum = -1
		}
		data := model.GetCateGory(pageSize, pageNum)
		code = errmsg.SUCCESS
		c.JSON(http.StatusOK, gin.H{
			"status": code,
			"data": data,
			"message": errmsg.GetErrMsg(code),
		})
}
// 编辑分类
func EditCategory(c *gin.Context)  {
	var data model.Category
	id, _ := strconv.Atoi(c.Param("id"))
	_ = c.ShouldBindJSON(&data)
	code = model.CheckCategory(data.Name)
	if code == errmsg.SUCCESS{
		model.EditCategory(id, &data)
	}
	if code == errmsg.ErrorCategoryUsed{
		c.Abort()
	}
	c.JSON(http.StatusOK, gin.H{
		"status": code,
		"message": errmsg.GetErrMsg(code),
	})
}

// 删除分类
func DeleteCategory(c *gin.Context)  {
	id, _ := strconv.Atoi(c.Param("id"))
	code = model.DeleteCategory(id)
	c.JSON(http.StatusOK, gin.H{
		"status": code,
		"message": errmsg.GetErrMsg(code),
	})
}




