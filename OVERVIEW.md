# OpenRouter Go SDK

The OpenRouter Go SDK is a type-safe client for building AI applications with access to 400+ language models through a unified API.

## Why use the OpenRouter SDK?

Integrating AI models into applications involves handling different provider APIs, managing model-specific requirements, and avoiding common implementation mistakes. The OpenRouter SDK standardizes these integrations with idiomatic Go types, retries, and typed errors.

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
						"Explain quantum computing",
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

The SDK provides three core benefits:

### Auto-generated from API specifications

The SDK is automatically generated from OpenRouter's OpenAPI specs and updated with every API change. New models, parameters, and features appear in generated types and docs immediately.

```go
res, err := s.Chat.Send(ctx, components.ChatRequest{
	Model: openrouter.Pointer("openai/gpt-4o"),
	// ...
}, nil)
```

### Type-safe by default

Request and response types live under `models/components` and `models/operations`. API errors are mapped to typed values under `models/sdkerrors`.

```go
res, err := s.Chat.Send(ctx, components.ChatRequest{
	Model: openrouter.Pointer("openai/gpt-4o"),
	Messages: []components.ChatMessages{
		components.CreateChatMessagesUser(
			components.ChatUserMessage{
				Role: components.ChatUserMessageRoleUser,
				Content: components.CreateChatUserMessageContentStr("Hello"),
			},
		),
	},
}, nil)
```

### Streaming and platform APIs

Use the same client for streaming chat completions, embeddings, rerank, TTS, video generation, and platform APIs such as API keys, credits, models, guardrails, and workspaces.

```go
res, err := s.Chat.Send(ctx, components.ChatRequest{
	Model:  openrouter.Pointer("openai/gpt-4o"),
	Messages: []components.ChatMessages{ /* ... */ },
	Stream:   openrouter.Pointer(true),
}, nil)
if res != nil && res.EventStream != nil {
	defer res.EventStream.Close()
	for res.EventStream.Next() {
		chunk := res.EventStream.Value()
		if chunk == nil {
			continue
		}
		for _, choice := range chunk.Data.Choices {
			if text, ok := choice.Delta.Content.Get(); ok && text != nil {
				_ = *text
			}
		}
	}
}
```

See [examples/chat-stream](examples/chat-stream) for a runnable streaming chat example.

## Installation

```bash
go get github.com/OpenRouterTeam/go-sdk
```

For beta releases, pin an explicit version:

```bash
go get github.com/OpenRouterTeam/go-sdk@v0.5.20
```

**Requirements:** Go 1.25 or higher

Get your API key from [openrouter.ai/settings/keys](https://openrouter.ai/settings/keys).

## Quick start

See [examples/README.md](examples/README.md) for runnable examples, starting with [examples/chat](examples/chat), or the [API Reference](https://openrouter.ai/docs/sdks/go/api-reference) for the full method list.
