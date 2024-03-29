package app

import (
	"context"
	"github.com/CuiYao631/DevMate/app/notify"
	"github.com/segmentio/ksuid"
	"github.com/wailsapp/wails/v2/pkg/menu"
	"github.com/wailsapp/wails/v2/pkg/runtime"
)

type Application struct {
	ctx          context.Context
	menus        *menu.Menu
	notification notify.Notify
}

func New() *Application {
	return &Application{}
}
func (a *Application) Startup(ctx context.Context, menus *menu.Menu) {
	a.ctx = ctx
	a.menus = menus

	//menus.Append(meu.NewWindowMenu(ctx))
	//runtime.MenuSetApplicationMenu(ctx, menus)
	//runtime.MenuUpdateApplicationMenu(ctx)

	//f := f ind.NewFind(ctx, "", "", "")
	//go f.OpenFind()
	//go f.OpenBeFound()
	n := notify.MakeNotify()
	a.notification = n

}
func (a *Application) ResizeMainWindow(width, height float64) {
	if width > 1000 {
		width = 1000
	}
	runtime.WindowSetSize(a.ctx, int(width), int(height))
}

func (a *Application) DefaultSizeMainWindow() {
	runtime.WindowSetSize(a.ctx, 350, 350)
}
func (a *Application) CreateKsuid() string {
	return ksuid.New().String()
}
func (a *Application) Notify() {
	a.notification.Notify("test", "test")
}
