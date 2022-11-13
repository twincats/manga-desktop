package main

import (
	"embed"

	"mangav4/system"
	"mangav4/system/file"
	"mangav4/system/manga"

	"github.com/wailsapp/wails/v2"
	"github.com/wailsapp/wails/v2/pkg/options"
	"github.com/wailsapp/wails/v2/pkg/options/assetserver"
	"github.com/wailsapp/wails/v2/pkg/options/windows"
)

//go:embed all:frontend/dist
var assets embed.FS

func main() {
	// Single instance check
	if _, err := system.CreateMutexW(system.ProcCreateMutexW, "mangav4.application"); err != nil {
		return
	}
	// Create an instance of the app structure
	app := NewApp()

	// Connect to database
	db := system.ConnectDatabasePostgres("mangav3")

	// Class bind instance
	manga := manga.NewManga(db)
	// reader := file.NewFilesDataReader()

	// Create application with options
	err := wails.Run(&options.App{
		Title:  "mangav4",
		Width:  1024,
		Height: 768,
		AssetServer: &assetserver.Options{
			Assets:  assets,
			Handler: file.NewFileLoader(),
		},
		Windows: &windows.Options{
			WebviewIsTransparent: false,
			WindowIsTranslucent:  false,
		},
		BackgroundColour: &options.RGBA{R: 27, G: 38, B: 54, A: 1},
		OnStartup:        app.startup,
		Bind: []interface{}{
			app,
			manga,
			// reader,
		},
	})

	if err != nil {
		println("Error:", err.Error())
	}
}
