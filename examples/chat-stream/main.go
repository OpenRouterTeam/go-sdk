package main

import (
	"context"
	"fmt"
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
		Model:  openrouter.Pointer("openai/gpt-4o-mini"),
		Stream: openrouter.Pointer(true),
		Messages: []components.ChatMessages{
			components.CreateChatMessagesUser(
				components.ChatUserMessage{
					Role: components.ChatUserMessageRoleUser,
					Content: components.CreateChatUserMessageContentStr(
						"Count from 1 to 5, one number per line.",
					),
				},
			),
		},
		Temperature: optionalnullable.From(openrouter.Pointer(0.7)),
	}, nil)
	if err != nil {
		log.Fatal(err)
	}
	if res == nil || res.EventStream == nil {
		log.Fatal("expected streaming response")
	}

	stream := res.EventStream
	defer stream.Close()

	for stream.Next() {
		chunk := stream.Value()
		if chunk == nil {
			continue
		}
		for _, choice := range chunk.Data.Choices {
			if text, ok := choice.Delta.Content.Get(); ok && text != nil {
				fmt.Print(*text)
			}
		}
	}
	fmt.Println()

	if err := stream.Err(); err != nil {
		log.Fatal(err)
	}
}
