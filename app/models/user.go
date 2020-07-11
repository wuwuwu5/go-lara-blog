package models

import "time"

type User struct {
	ID            int
	Name          string `json:"name" from:"name" validate:"required,max=20" label:"用户名"`
	Email         string `validate:"email" label:"邮箱"`
	EmailVerified time.Time
	Password      string `validate:"required,min=6,max=20" label:"密码"`
	PasswordConfirmation      string `validate:"required,min=6,max=20" label:"确认密码"`
	RememberToken string
	CreatedAt     time.Time
	UpdateAt      time.Time
}
