package controller

import (
	"lara-blog/helpers"
	"log"
	"net/http"
)

type Index struct {
}

// 首页
func (this *Index) Index(w http.ResponseWriter, request *http.Request) {
	err := helpers.View(w, "./views/index/index.html", nil)

	if err != nil {
		log.Fatal(err)
		return
	}
}

// 帮助页
func (this *Index) Help(w http.ResponseWriter, r *http.Request) {

	err := helpers.View(w, "./views/index/help.html", nil)

	if err != nil {
		log.Fatal(err)
		return
	}
}

// 关于我
func (this *Index) About(w http.ResponseWriter, r *http.Request) {

	err := helpers.View(w, "./views/index/about.html", nil)

	if err != nil {
		log.Fatal(err)
		return
	}
}
