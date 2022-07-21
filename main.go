package main

import (
	"embed"
	"github.com/wailsapp/wails/v2/pkg/options/mac"

	"github.com/wailsapp/wails/v2"
	"github.com/wailsapp/wails/v2/pkg/options"
)

//go:embed frontend/dist
var assets embed.FS

func main() {
	// Create an instance of the app structure
	app := NewApp()

	// Create application with options
	err := wails.Run(&options.App{
		Title:            "next",
		Width:            1024,
		Height:           768,
		MinWidth:         1024,
		MinHeight:        768,
		Assets:           assets,
		BackgroundColour: &options.RGBA{R: 27, G: 38, B: 54, A: 1},
		OnStartup:        app.startup,
		Bind: []interface{}{
			app,
		},
		Mac: &mac.Options{
			TitleBar:             mac.TitleBarHidden(),
			WebviewIsTransparent: true,
			About: &mac.AboutInfo{
				Title:   "next",
				Message: "©️ 2022 next",
				Icon:    nil,
			},
		},
	})

	if err != nil {
		println("Error:", err)
	}
}
