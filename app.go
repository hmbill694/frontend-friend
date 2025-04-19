package main

import (
	"context"
	"database/sql"
	"fmt"
	"frontend-friend/backend/database"
	actiongraphs "frontend-friend/backend/graph"
	"frontend-friend/backend/models"
	"frontend-friend/backend/utils"

	"github.com/tmc/langchaingo/llms"
	"github.com/tmc/langgraphgo/graph"
)

// App struct
type App struct {
	ctx                context.Context
	db                 *sql.DB
	htmlGeneratorGraph *graph.Runnable
}

// NewApp creates a new App application struct
func NewApp() *App {
	return &App{}
}

// startup is called when the app starts. The context is saved
// so we can call the runtime methods
func (a *App) startup(ctx context.Context) {
	a.ctx = ctx

	_, err := database.CreateDB()

	utils.PanicIfError(err)

	db, err := database.CreateConnection()

	utils.PanicIfError(err)

	a.db = db

	err = database.RunMigrations(a.db)

	utils.PanicIfError(err)

	geminiModel, err := models.NewGeminiModel("gemini-2.0-flash")

	utils.PanicIfError(err)

	htmlGraph, err := actiongraphs.CreateGraphHtmlGeneratorGraph(geminiModel)

	utils.PanicIfError(err)

	a.htmlGeneratorGraph = htmlGraph
}

func (a *App) shutdown(ctx context.Context) {
	a.db.Close()
}

// Greet returns a greeting for the given name
func (a *App) Greet(name string) string {
	return fmt.Sprintf("Hello %s, It's show time!", name)
}

func (a *App) GenerateHTMLPage(userPrompt string) bool {

	a.htmlGeneratorGraph.Invoke(a.ctx, []llms.MessageContent{
		llms.TextParts(llms.ChatMessageTypeHuman, "a todo app"),
	})

	return false
}
