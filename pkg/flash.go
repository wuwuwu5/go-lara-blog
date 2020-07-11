package pkg

import (
	"encoding/base64"
	"net/http"
)

//// cookie 闪存
//// 用于存储validate验证信息
//type FlashData struct {
//	Name string            // cookie储存的名称
//	Data map[string]string // 错误信息
//}
//
////// 实例化闪存
////func NewFlashData(name string) *FlashData {
////	return &FlashData{
////		Name: name,
////		Data: make(map[string]string),
////	}
////}

const (
	ValidateCookieName = "validateMessage"
)

// 保存验证消息闪存
func SaveValidateMessages(w http.ResponseWriter, message string) {
	cookie := http.Cookie{
		Name:   ValidateCookieName,
		Value:  base64.StdEncoding.EncodeToString([]byte(message)),
		Path:   "/",
		MaxAge: 100,
	}

	http.SetCookie(w, &cookie)
}

// 删除验证消息闪存
func DelValidateMessages(w http.ResponseWriter) {
	cookie := http.Cookie{
		Name:   ValidateCookieName,
		Value:  "",
		Path:   "/",
		MaxAge: -1,
	}

	http.SetCookie(w, &cookie)
}
