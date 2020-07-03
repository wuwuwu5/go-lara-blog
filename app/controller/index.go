package controller

import (
	"html/template"
	"log"
	"net/http"
)

type Index struct {
}

// 首页
func (this *Index) Index(writer http.ResponseWriter, request *http.Request) {
	tmpl, err := template.ParseFiles("./views/layout/app.html", "./views/index/index.html")

	if err != nil {
		log.Fatal(err)
		return
	}

	tmpl.Execute(writer, nil)
}

// 帮助页
func (this *Index) Help(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles("./views/layout/app.html", "./views/index/help.html")

	if err != nil {
		log.Fatal(err)
		return
	}

	tmpl.Execute(w, nil)
}

// 关于我
func (this *Index) About(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles("./views/layout/app.html", "./views/index/about.html")

	if err != nil {
		log.Fatal(err)
		return
	}

	tmpl.Execute(w, nil)
}
