package routes

import (
	v1 "ginblog1/api/v1"
	"ginblog1/middleware"
	"ginblog1/utils"
	"github.com/gin-gonic/gin"
)

func InitRouter() {
	gin.SetMode(utils.AppMode)
	//r := gin.Default()
	r := gin.New()
	r.Use(gin.Recovery())
	r.Use(middleware.Logger())
	r.Use(middleware.Cors())
	authRouter := r.Group("api/v1")
	authRouter.Use(middleware.JwtToken())
	{
		// User模块的路由接口
		authRouter.POST("user/add", v1.AddUser)
		authRouter.PUT("user/:id", v1.EditUser)
		authRouter.DELETE("user/:id", v1.DeleteUser)
		authRouter.GET("users", v1.GetUsers)
		// Category模块的路由接口
		authRouter.POST("category/add", v1.AddCategory)
		authRouter.PUT("category/:id", v1.EditCategory)
		authRouter.DELETE("category/:id", v1.DeleteCategory)
		authRouter.GET("categorys", v1.GetCategorys)
		authRouter.GET("category/:id", v1.GetCategoryAllArticle)
		// Article模块的路由接口
		authRouter.POST("article/add", v1.AddArticle)
		authRouter.PUT("article/:id", v1.EditArticle)
		authRouter.DELETE("article/:id", v1.DeleteArticle)
		authRouter.GET("articles", v1.GetArticles)
		authRouter.GET("article/:id", v1.GetArticle)
		// Student模块的路由接口
		authRouter.POST("student/add", v1.AddStudent)
		authRouter.GET("student/:id", v1.GetStudentBooks)
		// Book模块的路由接口
		authRouter.POST("book/add", v1.AddBook)
	}
	publicRouter := r.Group("/api/v1")
	{
		publicRouter.POST("login", v1.Login)
	}
	r.Run(utils.HttpPort)
}