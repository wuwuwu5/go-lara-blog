package controller

import (
	"github.com/go-playground/validator/v10"
	"lara-blog/app/models"
	"lara-blog/helpers"
	"lara-blog/pkg"
	"lara-blog/routes"
	"net/http"
)

type User struct {
}

func (this *User) Signup(c *routes.Context) {
	helpers.View(c, "./views/users/create.html", routes.H{})
}

// 账号注册
func (this *User) Register(c *routes.Context) {
	c.Req.ParseForm()

	user := &models.User{
		Name:                 c.PostForm("name"),
		Email:                c.PostForm("email"),
		Password:             c.PostForm("password"),
		PasswordConfirmation: c.PostForm("password_confirmation"),
	}

	err := pkg.Validate.Struct(user)

	if err != nil {
		pkg.SaveValidateMessages(c.Writer, pkg.Translate(err.(validator.ValidationErrors)))

		// 返回页面
		http.Redirect(c.Writer, c.Req, "/signup", 302)
	}
}
