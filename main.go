package main

import (
	. "lara-blog/app/controller"
	"lara-blog/app/models"
	"net/http"
)

func main() {

	// 初始化数据库连接
	models.InitDB()

	// 静态文件
	// StripPrefix 替换/static为/public
	http.Handle("/static/", http.StripPrefix("/static", http.FileServer(http.Dir("./public"))))

	indexController := &Index{}
	http.HandleFunc("/", indexController.Index)
	http.HandleFunc("/help", indexController.Help)
	http.HandleFunc("/about", indexController.About)


	http.ListenAndServe(":8888", nil)
}
