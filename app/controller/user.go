package controller

import (
	"fmt"
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
		return
	}

	// 验证name 和 邮箱一致性
	err = user.CheckUserName()

	if err != nil {
		pkg.SaveValidateMessages(c.Writer, "用户名重复")

		fmt.Println(err)

		// 返回页面
		http.Redirect(c.Writer, c.Req, "/signup", 302)
	}

	// 验证name 和 邮箱一致性
	err = user.CheckUserEmail()

	if err != nil {
		pkg.SaveValidateMessages(c.Writer, "邮箱重复")

		fmt.Println(err)

		// 返回页面
		http.Redirect(c.Writer, c.Req, "/signup", 302)
	}

	if err != nil {
		pkg.SaveValidateMessages(c.Writer, "邮箱重复")

		fmt.Println(err)

		// 返回页面
		http.Redirect(c.Writer, c.Req, "/signup", 302)
	}

	// 添加数据
	err = user.CreateUser()

	if err != nil {
		pkg.SaveValidateMessages(c.Writer, "添加失败")

		fmt.Println(err)

		// 返回页面
		http.Redirect(c.Writer, c.Req, "/signup", 302)
	}

	http.Redirect(c.Writer, c.Req, "/", 302)
}
