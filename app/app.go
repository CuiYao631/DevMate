package app

import (
	"context"
	"github.com/segmentio/ksuid"
	"github.com/wailsapp/wails/v2/pkg/runtime"
)

type Application struct {
	ctx context.Context
}

func New() *Application {
	return &Application{}
}

func (a *Application) Startup(ctx context.Context) {
	a.ctx = ctx
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
