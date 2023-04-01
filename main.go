package main

import (
	"embed"

	"github.com/wailsapp/wails/v2"
	"github.com/wailsapp/wails/v2/pkg/options"
	"github.com/wailsapp/wails/v2/pkg/options/assetserver"
)

//go:embed all:frontend/dist
var assets embed.FS

func main() {
	App := NewApp()

	err := wails.Run(&options.App{
		Title:  "file-organizer-desktop",
		Width:  700,
		Height: 600,
		AssetServer: &assetserver.Options{
			Assets: assets,
		},
		OnStartup: App.startup,
		Bind: []interface{}{
			App,
		},
	})

	if err != nil {
		println("Error:", err.Error())
	}
}
