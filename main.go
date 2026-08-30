package main

import (
	"context"
	"embed"

	"github.com/wailsapp/wails/v2"
	"github.com/wailsapp/wails/v2/pkg/options"
	"github.com/wailsapp/wails/v2/pkg/options/assetserver"
	"github.com/wailsapp/wails/v2/pkg/options/mac"
	"github.com/wailsapp/wails/v2/pkg/runtime"
)

//go:embed all:frontend/dist
var assets embed.FS

func main() {
	app := NewApp()

	err := wails.Run(&options.App{
		Title:     "Sift",
		Width:     1280,
		Height:    800,
		MinWidth:  1000,
		MinHeight: 600,

		AssetServer: &assetserver.Options{
			Assets: assets,
		},

		BackgroundColour: &options.RGBA{
			R: 15,
			G: 17,
			B: 21,
			A: 255,
		},

		OnStartup:  app.startup,
		OnShutdown: app.shutdown,

		Bind: []interface{}{
			app,
		},

		Mac: &mac.Options{
			TitleBar: &mac.TitleBar{
				TitlebarAppearsTransparent: true,
				HideTitle:                  false,
			},
			WebviewIsTransparent: false,
			WindowIsTranslucent:  false,
			About: &mac.AboutInfo{
				Title:   "Sift",
				Message: "Open-source SQL client",
			},
		},

		OnDomReady: func(ctx context.Context) {
			runtime.WindowSetTitle(ctx, "Sift")
		},
	})

	if err != nil {
		println("Error:", err.Error())
	}
}
