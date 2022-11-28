package main

import (
	"context"
	"fmt"

	"mangav4/system/app"

	"github.com/wailsapp/wails/v2/pkg/runtime"
)

// App struct
type App struct {
	ctx context.Context
}

// NewApp creates a new App application struct
func NewApp() *App {
	return &App{}
}

// startup is called when the app starts. The context is saved
// so we can call the runtime methods
func (a *App) startup(ctx context.Context) {
	a.ctx = ctx
	go app.Startup(&ctx)
}

func (b *App) beforeClose(ctx context.Context) (prevent bool) {
	if app.BlockClose {
		dialog, err := runtime.MessageDialog(ctx, runtime.MessageDialogOptions{
			Type:    runtime.QuestionDialog,
			Title:   "Quit?",
			Message: "Are you sure you want to quit?",
		})

		if err != nil {
			return false
		}
		return dialog != "Yes"
	}
	return false
}

// Greet returns a greeting for the given name
func (a *App) Greet(name string) string {
	return fmt.Sprintf("Hello %s, It's show time!", name)
}
