//go:build !darwin
// +build !darwin

package notify

import (
	"github.com/go-toast/toast"
	"log"
)

func NotifyMessage(appId, title, message string) {

	notification := toast.Notification{

		AppID:   appId,
		Title:   title,
		Message: message,
		Icon:    "D:\\logo.png", // 文件必须存在
		//Actions: []toast.Action{
		//	{"protocol", "按钮1", "https://www.google.com/"},
		//	{"protocol", "按钮2", "https://github.com/"},
		//},
		//Audio: toast.Mail,
	}
	err := notification.Push()
	if err != nil {
		log.Println(err)
	}
}
