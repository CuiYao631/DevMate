package main

import (
	"context"
	"embed"
	"github.com/CuiYao631/DevMate/app"
	"github.com/wailsapp/wails/v2"
	"github.com/wailsapp/wails/v2/pkg/menu"
	"github.com/wailsapp/wails/v2/pkg/options"
	"github.com/wailsapp/wails/v2/pkg/options/assetserver"
	"github.com/wailsapp/wails/v2/pkg/options/mac"
)

//go:embed all:frontend/dist
var assets embed.FS

//go:embed build/appicon.png
var icon []byte

func main() {
	// Create an instance of the app structure
	application := app.New()
	// 创建默认菜单
	defaultMenu := menu.NewMenu()
	//// 添加菜单
	//defaultMenu.Append(menu.AppMenu())
	//// 添加编辑菜单
	//defaultMenu.Append(menu.EditMenu())

	// Create application with options
	err := wails.Run(&options.App{
		Title:  "DevMate",
		Width:  800,
		Height: 600,
		AssetServer: &assetserver.Options{
			Assets: assets,
		},
		//BackgroundColour: &options.RGBA{R: 27, G: 38, B: 54, A: 1},
		OnStartup: func(ctx context.Context) {
			application.Startup(ctx, defaultMenu)
		},
		Bind: []interface{}{
			application,
		},
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
		//Windows: &windows.Options{},
	})

	if err != nil {
		println("Error:", err.Error())
	}
}
