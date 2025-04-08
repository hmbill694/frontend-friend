package main

import (
	"embed"
	"os"
	"strings"

	"github.com/pocketbase/pocketbase"
	"github.com/pocketbase/pocketbase/cmd"
	"github.com/wailsapp/wails/v2"
	"github.com/wailsapp/wails/v2/pkg/options"
	"github.com/wailsapp/wails/v2/pkg/options/assetserver"

	"github.com/pocketbase/pocketbase/plugins/migratecmd"
	// enable once you have at least one migration
	_ "frontend-friend/migrations"
)

//go:embed all:frontend/dist
var assets embed.FS

func main() {
	// Create an instance of the app structure
	app := NewApp()

	// Go Routine to run Pocketbase
	go func() {
		pocketBaseApp := pocketbase.New()

		// loosely check if it was executed using "go run"
		isGoRun := strings.HasPrefix(os.Args[0], os.TempDir())

		migratecmd.MustRegister(pocketBaseApp, pocketBaseApp.RootCmd, migratecmd.Config{
			// enable auto creation of migration files when making collection changes in the Dashboard
			// (the isGoRun check is to enable it only during development)
			Automigrate: isGoRun,
		})

		pocketBaseApp.Bootstrap()
		serveCmd := cmd.NewServeCommand(pocketBaseApp, true)
		serveCmd.Execute()
	}()

	// Create application with options
	err := wails.Run(&options.App{
		Title:            "frontend-friend",
		WindowStartState: options.WindowStartState(3),
		AssetServer: &assetserver.Options{
			Assets: assets,
		},
		BackgroundColour: &options.RGBA{R: 27, G: 38, B: 54, A: 1},
		OnStartup:        app.startup,
		Bind: []interface{}{
			app,
		},
	})

	if err != nil {
		println("Error:", err.Error())
	}
}
