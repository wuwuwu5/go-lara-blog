package controller

import (
	"lara-blog/helpers"
	"lara-blog/routes"
)

type Index struct {
}

// 首页
func (this *Index) Index(c *routes.Context) {
	helpers.View(c, "./views/index/index.html", nil);
}

// 帮助页
func (this *Index) Help(c *routes.Context) {
	helpers.View(c, "./views/index/help.html", nil);
}

// 关于我
func (this *Index) About(c *routes.Context) {
	helpers.View(c, "./views/index/about.html", nil);
}
