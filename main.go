package main

import (
	. "lara-blog/app/controller"
	"lara-blog/app/models"
	_ "lara-blog/pkg"
	"lara-blog/routes"
)

func main() {

	// 初始化数据库连接
	models.InitDB()

	// 启动路由
	router := routes.InitRouter()

	indexController := &Index{}
	{
		router.Get("/", indexController.Index);
		router.Get("/help", indexController.Help);
		router.Get("/about", indexController.About);
	}

	// 用户
	userController := &User{}
	{
		router.GET("/signup", userController.Signup)
		router.POST("/signup", userController.Register)
	}

	router.Run(":8888")
}
