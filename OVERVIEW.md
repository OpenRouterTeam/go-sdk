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
	})
	if err != nil {
		log.Fatal(err)
	}
	if res != nil {
		defer res.Object.Close()
		for res.Object.Next() {
			log.Println(res.Object.Value())
		}
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
})
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
})
```

### Streaming and platform APIs

Use the same client for streaming chat completions, embeddings, rerank, TTS, video generation, and platform APIs such as API keys, credits, models, guardrails, and workspaces.

```go
res, err := s.Chat.Send(ctx, components.ChatRequest{
	Model: openrouter.Pointer("openai/gpt-4o"),
	Messages: []components.ChatMessages{ /* ... */ },
})
if res != nil {
	defer res.Object.Close()
	for res.Object.Next() {
		event := res.Object.Value()
		_ = event
	}
}
```

## Installation

```bash
go get github.com/OpenRouterTeam/go-sdk
```

For beta releases, pin an explicit version:

```bash
go get github.com/OpenRouterTeam/go-sdk@v0.5.0
```

**Requirements:** Go 1.22 or higher

Get your API key from [openrouter.ai/settings/keys](https://openrouter.ai/settings/keys).

## Quick start

See [examples/chat](examples/chat) for a runnable example, or the [API Reference](https://openrouter.ai/docs/sdks/go/api-reference) for the full method list.
