package main

import (
	"context"
	"embed"
	"github.com/CuiYao631/DevMate/app"
	"github.com/wailsapp/wails/v2/pkg/menu"
	"github.com/wailsapp/wails/v2/pkg/options/mac"

	"github.com/wailsapp/wails/v2"
	"github.com/wailsapp/wails/v2/pkg/options"
	"github.com/wailsapp/wails/v2/pkg/options/assetserver"
)

//go:embed all:frontend/dist
var assets embed.FS

//go:embed build/appicon.png
var icon []byte

func main() {
	application := app.New()

	// 创建默认菜单
	defaultMenu := menu.NewMenu()
	// 添加菜单
	defaultMenu.Append(menu.AppMenu())
	// 添加编辑菜单
	defaultMenu.Append(menu.EditMenu())

	err := wails.Run(&options.App{
		Title:  "DevMate",
		Width:  800,
		Height: 600,
		AssetServer: &assetserver.Options{
			Assets: assets,
		},
		HideWindowOnClose: true,
		//BackgroundColour:  &options.RGBA{R: 255, G: 255, B: 255, A: 0},
		OnStartup: func(ctx context.Context) {
			application.Startup(ctx, defaultMenu)
		},
		Bind: []interface{}{
			application,
		},
		//Frameless:       true,
		//CSSDragProperty: "window",
		//CSSDragValue:    "1",
		Mac: &mac.Options{
			TitleBar:             mac.TitleBarHiddenInset(),
			Appearance:           mac.NSAppearanceNameDarkAqua,
			WebviewIsTransparent: true,
			WindowIsTranslucent:  true,
			About: &mac.AboutInfo{
				Title:   "DevMate",
				Message: "© 2024 Alex",
				Icon:    icon,
			},
		},
	})

	if err != nil {
		println("Error:", err.Error())
	}
}
