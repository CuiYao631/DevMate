package menu

import (
	"context"
	"fmt"
	"github.com/wailsapp/wails/v2/pkg/menu"
	"github.com/wailsapp/wails/v2/pkg/runtime"
)

var (
	// AlwaysOnTop 总是在最前
	AlwaysOnTop = false
)

func NewWindowMenu(ctx context.Context) *menu.MenuItem {
	wm := NewMenu("window")
	text := wm.AddText("")
	update := func() {
		text.Label = fmt.Sprintf("Always On Top: %v", AlwaysOnTop)
		runtime.MenuUpdateApplicationMenu(ctx)
	}
	wm.AddSeparator()
	wm.Add("Reload", func(data *menu.CallbackData) {
		runtime.WindowReload(ctx)
	})

	wm.Add("Reload App", func(data *menu.CallbackData) {
		runtime.WindowReloadApp(ctx)
	})
	wm.AddSeparator()
	wm.Add("Default", func(data *menu.CallbackData) {

	})

	wm.Add("Maximise", func(data *menu.CallbackData) {
		runtime.WindowToggleMaximise(ctx)
	})
	wm.Add("Minimise", func(data *menu.CallbackData) {
		runtime.WindowMinimise(ctx)
	})
	wm.AddSeparator()
	wm.Add("Always On Top", func(data *menu.CallbackData) {
		AlwaysOnTop = !AlwaysOnTop
		runtime.WindowSetAlwaysOnTop(ctx, AlwaysOnTop)
		update()
	})
	wm.Add("Move To Center", func(data *menu.CallbackData) {
		runtime.WindowCenter(ctx)
	})
	wm.Add("Message Box", func(data *menu.CallbackData) {
		runtime.MessageDialog(ctx, runtime.MessageDialogOptions{
			Type:          runtime.QuestionDialog,
			Title:         "提示",
			Message:       "这是一个消息框",
			Buttons:       []string{"确定", "取消"},
			DefaultButton: "确定",
			CancelButton:  "确定",
			//Icon:          [],
		})
	})
	update()
	return wm.Build()
}
