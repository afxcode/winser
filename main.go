package main

import (
	"embed"

	"github.com/wailsapp/wails/v2"
	"github.com/wailsapp/wails/v2/pkg/options"
)

//go:embed all:frontend/dist
var assets embed.FS

func main() {
	// Create an instance of the app structure
	app := NewApp()

	// Create application with options
	err := wails.Run(&options.App{
		Title:     "Winser - Windows Service Manager",
		Width:     600,
		Height:    370,
		MinWidth:  600,
		MinHeight: 370,
		MaxWidth:  600,
		MaxHeight: 370,
		Assets:    assets,
		// BackgroundColour: &options.RGBA{R: 255, G: 255, B: 255, A: 0},
		OnStartup: app.startup,
		Bind: []any{
			app,
		},
	})

	if err != nil {
		println("Error:", err.Error())
	}
}
