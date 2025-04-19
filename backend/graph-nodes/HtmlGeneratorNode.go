package graphnodes

import (
	"context"
	"fmt"
	"frontend-friend/backend/models"
	"strings"

	"github.com/tmc/langchaingo/llms"
)

type PageCreatorGraph struct {
	Messages       []any
	HTMLIsValid    bool
	UserInput      string
	OutputHTML     *string
	OutputFileName *string
}

type DefaultGraphState any

type GraphStepFunction[GS DefaultGraphState] func(state GS) (map[string]interface{}, error)

type ChatMessage struct {
	Content string
	Role    string
}

func HTMLGenerator(llm models.ChatModel) func(ctx context.Context, state []llms.MessageContent) ([]llms.MessageContent, error) {
	return func(ctx context.Context, state []llms.MessageContent) ([]llms.MessageContent, error) {
		userInput := state[len(state)-1]

		prompt := fmt.Sprintf(`
			Act as an expert software developer and produce the html on demand.

			Output only the html to accomplish this task. Do not output anything other than valid html pages.

			Do not make up any invalid html. Know that your work is critical to the success of your team and
			that the generation of correct, well formatted html will be rewarded handsomely with lots of acceptance
			and happiness. Ensure that you do not surround your HTML with anything. We are not rendering the HTML in
			a markdown file. Everyone will thank you kindly should you ensure you do this. It is your destiny to do this correctly.
			Your team has very strict styling and coding standards. Your preformance will be evaluated
			very negatively should you deviate from the teams standards. The teams standards are as
			follows.

			- All styling should be done using TailwindCSS.
			- All generated html documents should include the include script for TailwindCSS. You team is very interested in no build html.
			- All generated pages should be responsive. Being able to adjust to mobile, tablet and desktop.
			- All generated html documents should not have any javascript within them that is not related to styling. Your team will add any JS required to make the page interactive.
			- All generated html documents should be be accessible.

			Below is the description of the web page that we need the html for. Take a deep breath you got this.
			%s
		`, userInput)

		msgContent, err := llm.Call(context.Background(), prompt)
		if err != nil {
			return nil, fmt.Errorf("failed to invoke LLM: %w", err)
		}

		trimmedMsg := strings.ReplaceAll(msgContent, "```html", "")
		trimmedMsg = strings.ReplaceAll(trimmedMsg, "```", "")

		return append(state,
			llms.TextParts(llms.ChatMessageTypeAI, trimmedMsg),
		), nil
	}
}
