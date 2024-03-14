//go:build darwin
// +build darwin

package notify

import (
	gosxnotifier "github.com/deckarep/gosx-notifier"
	"log"
)

func NotifyMessage(title, message string) {
	//至少指定要显示给最终用户的消息.
	note := gosxnotifier.NewNotification("Check your Apple Stock!")

	//(可选) 设置标题
	note.Title = "这是一个弹窗标题"

	//(可选) 设置字幕
	note.Subtitle = "我的字幕"

	//(可选) 从预定义的集合中设置声音.
	note.Sound = gosxnotifier.Default

	//(可选) 设置一个组，确保只显示一个通知，替换同一组id的先前通知.
	//note.Group = "com.unique.yourapp.identifier"

	//(可选) 设置发件人 (通知现在将使用Safari图标)
	note.Sender = "com.wails.DevMate"

	//(可选) 指定要在单击通知时打开的url或bundleid。
	note.Link = "www.baidu.com" //or BundleID like: com.apple.Terminal

	//(可选) 应用程序图标 (仅限10.9)
	//note.AppIcon = "build/windows/icon.ico"

	//(可选) 内容图像 (仅限10.9)
	note.ContentImage = "build/appicon.png"

	//然后，推送通知
	err := note.Push()

	//检查错误
	if err != nil {
		log.Println(err)
		log.Println("Uh oh!")
	}
}
