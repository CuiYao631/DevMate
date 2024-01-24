package app

import (
	"context"
	"github.com/CuiYao631/DevMate/app/chatMessage/find"
	meu "github.com/CuiYao631/DevMate/app/menu"
	"github.com/segmentio/ksuid"
	"github.com/wailsapp/wails/v2/pkg/menu"
	"github.com/wailsapp/wails/v2/pkg/runtime"
	"log"
)

type Application struct {
	ctx   context.Context
	menus *menu.Menu
}

func New() *Application {
	return &Application{}
}

func (a *Application) Startup(ctx context.Context, menus *menu.Menu) {
	a.ctx = ctx
	a.menus = menus

	menus.Append(meu.NewWindowMenu(ctx))
	runtime.MenuSetApplicationMenu(ctx, menus)
	runtime.MenuUpdateApplicationMenu(ctx)

	f := find.NewFind(ctx, "", "", "")
	go f.OpenFind()
	go f.OpenBeFound()

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

// InitApp 初始化应用
func (a *Application) InitApp() {

}

func (a *Application) PushDataToReact() {
	data := "Hello from Go to React!"
	runtime.EventsEmit(a.ctx, "dataUpdated", data)
	log.Println("dataUpdated")
	//a.frontend.Events.Emit("dataUpdated", data)
}
