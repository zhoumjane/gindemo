package main

import (
	"fmt"
	"ginblog1/dao"
	"ginblog1/routes"
)

func main()  {
	_, err := dao.InitDb()
	if err != nil{
		fmt.Printf("连接数据库失败，请检查参数: %s", err)
		panic("mysql connect failed!")
	}
	routes.InitRouter()
}

