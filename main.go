package main

import (
	"embed"

	"github.com/wailsapp/wails/v2"
	"github.com/wailsapp/wails/v2/pkg/options"
)

//go:embed all:frontend/dist
var assets embed.FS

const version = "v0.3.0"

func main() {
	// read stored pined service name
	initPinedServiceNames()

	// Create an instance of the app structure
	app := NewApp()

	// Create application with options
	err := wails.Run(&options.App{
		Title:     "Winser - Windows Service Manager - " + version,
		Width:     720,
		Height:    370,
		MinWidth:  720,
		MinHeight: 370,
		MaxWidth:  0,
		MaxHeight: 0,
		Assets:    assets,
		OnStartup: app.startup,
		Bind: []any{
			app,
		},
	})

	if err != nil {
		println("Error:", err.Error())
	}
}
