package models

import (
	"crypto/md5"
	"fmt"
	"time"
)

type User struct {
	ID                   int
	Name                 string `json:"name" from:"name" validate:"required,max=20" label:"用户名"`
	Email                string `validate:"email" label:"邮箱"`
	EmailVerified        string
	Password             string `validate:"required,min=6,max=20" label:"密码"`
	PasswordConfirmation string `validate:"required,min=6,max=20" label:"确认密码"`
	RememberToken        string
	CreatedAt            string
	UpdateAt             string
}

// 验证用户名
func (this *User) CheckUserName() error {

	stmt, err := GetDB().Prepare("SELECT COUNT(*) FROM  users where name = ?")

	if err != nil {
		return err
	}

	row := stmt.QueryRow(this.Name)

	if row == nil {
		return nil
	}

	return fmt.Errorf("%s", "用户名已经存在")
}

// 验证用户邮箱
func (this *User) CheckUserEmail() error {
	stmt, err := GetDB().Prepare("select count(*) from  users where email = ?")

	if err != nil {
		return err
	}

	row := stmt.QueryRow(this.Email)

	if row == nil {
		return nil
	}

	return fmt.Errorf("%s", "邮箱已经存在")
}

// 创建用户
func (this *User) CreateUser() error {
	this.CreatedAt = time.Now().Format("2006-01-02 15:04:05")
	this.UpdateAt = time.Now().Format("2006-01-02 15:04:05")
	this.Password = fmt.Sprintf("%s", md5.Sum([]byte(this.Password)))

	stmt, err := GetDB().Prepare("insert  into users (`name`, `email`, `password`, `created_at`, `updated_at`) values(?,?,?,?,?)")

	if err != nil {
		return err
	}

	result, err := stmt.Exec(this.Name, this.Email, this.Password, this.CreatedAt, this.UpdateAt)

	if err != nil {
		return err
	}

	_, err = result.RowsAffected()

	if err != nil {
		return err
	}

	return nil
}
