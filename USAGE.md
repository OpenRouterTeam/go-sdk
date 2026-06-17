<!-- Start SDK Example Usage [usage] -->
```go
package main

import (
	"context"
	"log"
	"os"

	openrouter "github.com/OpenRouterTeam/go-sdk"
	"github.com/OpenRouterTeam/go-sdk/models/components"
	"github.com/OpenRouterTeam/go-sdk/optionalnullable"
)

func main() {
	ctx := context.Background()

	s := openrouter.New(
		openrouter.WithSecurity(os.Getenv("OPENROUTER_API_KEY")),
	)

	res, err := s.Chat.Send(ctx, components.ChatRequest{
		Model: openrouter.Pointer("openai/gpt-4o"),
		Messages: []components.ChatMessages{
			components.CreateChatMessagesUser(
				components.ChatUserMessage{
					Role: components.ChatUserMessageRoleUser,
					Content: components.CreateChatUserMessageContentStr(
						"Hello, how are you?",
					),
				},
			),
		},
		Temperature: optionalnullable.From(openrouter.Pointer(0.7)),
	}, nil)
	if err != nil {
		log.Fatal(err)
	}
	if res != nil && res.ChatResult != nil {
		log.Println(res.ChatResult.Choices)
	}
}

```
<!-- End SDK Example Usage [usage] -->