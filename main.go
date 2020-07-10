package main

import (
	. "lara-blog/app/controller"
	"lara-blog/routes"
)

func main() {

	// 初始化数据库连接
	//models.InitDB()

	// 启动路由
	router := routes.InitRouter()

	indexController := &Index{}
	{
		router.Get("/", indexController.Index);
		router.Get("/help", indexController.Help);
		router.Get("/about", indexController.About);
	}


	router.Run(":8888")
}
