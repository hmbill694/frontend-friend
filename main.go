package main

import (
	"embed"

	"github.com/wailsapp/wails/v2"
	"github.com/wailsapp/wails/v2/pkg/options"
	"github.com/wailsapp/wails/v2/pkg/options/assetserver"

	"github.com/joho/godotenv"
	_ "github.com/tursodatabase/go-libsql"
)

//go:embed all:frontend/dist
var assets embed.FS

func main() {
	err := godotenv.Load()
	app := NewApp()

	// Create application with options
	err = wails.Run(&options.App{
		Title:            "frontend-friend",
		WindowStartState: options.WindowStartState(3),
		AssetServer: &assetserver.Options{
			Assets: assets,
		},
		BackgroundColour: &options.RGBA{R: 27, G: 38, B: 54, A: 1},
		OnStartup:        app.startup,
		OnShutdown:       app.shutdown,
		// Expose the app to the frontend
		Bind: []any{
			app,
		},
	})

	if err != nil {
		panic("Error: " + err.Error())
	}
}
