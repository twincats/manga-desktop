package main

import (
	"embed"

	"mangav4/system"
	"mangav4/system/file"

	vips "github.com/twincats/golibvips/libvips"

	"github.com/wailsapp/wails/v2"
	"github.com/wailsapp/wails/v2/pkg/options"
	"github.com/wailsapp/wails/v2/pkg/options/assetserver"
	"github.com/wailsapp/wails/v2/pkg/options/windows"
)

//go:embed all:frontend/dist
var assets embed.FS

func main() {
	// Create an instance of the app structure
	mangaApp := NewApp()

	// Connect to database
	system.DatabaseStartUp()

	// Starting vips & shutdown after finish used
	vips.DisableConsoleLogging()
	vips.Startup(nil)
	defer vips.Shutdown()

	// Class bind instance
	binding := system.InitializeBinding()
	binding = append(binding, mangaApp)

	// Create application with options
	err := wails.Run(&options.App{
		Title:     "mangav4",
		Width:     1024,
		Height:    768,
		MinWidth:  1024,
		MinHeight: 768,
		// WindowStartState: options.Maximised,
		AssetServer: &assetserver.Options{
			Assets:  assets,
			Handler: file.NewFileLoader(),
		},
		Windows: &windows.Options{
			WebviewIsTransparent: false,
			WindowIsTranslucent:  false,
		},
		BackgroundColour: &options.RGBA{R: 24, G: 24, B: 24, A: 1},
		SingleInstanceLock: &options.SingleInstanceLock{
			UniqueId:               "a620c3a7-c432-4b96-93c0-d1f935ed3bff",
			OnSecondInstanceLaunch: mangaApp.onSecondInstanceLaunch,
		},
		OnStartup:     mangaApp.startup,
		OnBeforeClose: mangaApp.beforeClose,
		Bind:          binding,
	})

	if err != nil {
		println("Error:", err.Error())
	}
}
