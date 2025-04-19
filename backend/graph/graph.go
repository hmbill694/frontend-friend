package actiongraphs

import (
	"context"
	graphnodes "frontend-friend/backend/graph-nodes"
	"frontend-friend/backend/models"

	"github.com/tmc/langchaingo/llms"
	"github.com/tmc/langgraphgo/graph"
)

func CreateGraphHtmlGeneratorGraph(model models.ChatModel) (*graph.Runnable, error) {
	g := graph.NewMessageGraph()

	g.AddNode("html_generator", graphnodes.HTMLGenerator(model))

	g.AddNode(graph.END, func(ctx context.Context, state []llms.MessageContent) ([]llms.MessageContent, error) {
		return state, nil
	})

	g.AddEdge("oracle", graph.END)
	g.SetEntryPoint("oracle")

	return g.Compile()
}
