package openrouter_test

import (
	"context"
	"fmt"
	"log"
	"os"

	openrouter "github.com/OpenRouterTeam/go-sdk"
	"github.com/OpenRouterTeam/go-sdk/models/components"
	"github.com/OpenRouterTeam/go-sdk/models/operations"
	"github.com/OpenRouterTeam/go-sdk/optionalnullable"
)

// Example demonstrates basic usage of the OpenRouter SDK for chat completions.
// This example shows how to create a client, configure authentication, and send a chat request.
func Example() {
	ctx := context.Background()

	// Initialize the SDK with your API key
	sdk := openrouter.New(
		openrouter.WithSecurity(os.Getenv("OPENROUTER_API_KEY")),
	)

	// Send a chat completion request
	res, err := sdk.Chat.Send(ctx, components.ChatRequest{
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

	if res != nil && res.ChatResult != nil && len(res.ChatResult.Choices) > 0 {
		fmt.Printf("Response received with %d choices\n", len(res.ChatResult.Choices))
	}
}

// Example_chatWithMaxTokens demonstrates sending a chat request with max tokens limit.
func Example_chatWithMaxTokens() {
	ctx := context.Background()
	sdk := openrouter.New(
		openrouter.WithSecurity(os.Getenv("OPENROUTER_API_KEY")),
	)

	res, err := sdk.Chat.Send(ctx, components.ChatRequest{
		Model: openrouter.Pointer("anthropic/claude-3-sonnet"),
		Messages: []components.ChatMessages{
			components.CreateChatMessagesUser(
				components.ChatUserMessage{
					Role:    components.ChatUserMessageRoleUser,
					Content: components.CreateChatUserMessageContentStr("Say hello"),
				},
			),
		},
		MaxTokens: optionalnullable.From(openrouter.Pointer(int64(100))),
	}, nil)
	if err != nil {
		log.Fatal(err)
	}

	if res != nil && res.ChatResult != nil {
		fmt.Printf("Model used: %s\n", res.ChatResult.Model)
	}
}

// Example_listModels demonstrates listing available models.
func Example_listModels() {
	ctx := context.Background()
	sdk := openrouter.New(
		openrouter.WithSecurity(os.Getenv("OPENROUTER_API_KEY")),
	)

	res, err := sdk.Models.List(ctx, nil)
	if err != nil {
		log.Fatal(err)
	}

	if res != nil && len(res.Data) > 0 {
		fmt.Printf("Found %d models\n", len(res.Data))
		fmt.Printf("First model: %s\n", res.Data[0].Name)
	}
}

// Example_generateEmbedding demonstrates generating text embeddings.
func Example_generateEmbedding() {
	ctx := context.Background()
	sdk := openrouter.New(
		openrouter.WithSecurity(os.Getenv("OPENROUTER_API_KEY")),
	)

	res, err := sdk.Embeddings.Generate(ctx, operations.CreateEmbeddingsRequest{
		Model: "openai/text-embedding-ada-002",
		Input: operations.CreateInputUnionStr("The quick brown fox jumps over the lazy dog"),
	}, nil)
	if err != nil {
		log.Fatal(err)
	}

	if res != nil && res.CreateEmbeddingsResponseBody != nil && res.CreateEmbeddingsResponseBody.Data != nil {
		fmt.Printf("Embedding dimensions: %d\n", len(res.CreateEmbeddingsResponseBody.Data))
	}
}

// Example_getModel demonstrates retrieving information about a specific model.
func Example_getModel() {
	ctx := context.Background()
	sdk := openrouter.New(
		openrouter.WithSecurity(os.Getenv("OPENROUTER_API_KEY")),
	)

	res, err := sdk.Models.Get(ctx, "openai", "gpt-4o")
	if err != nil {
		log.Fatal(err)
	}

	if res != nil {
		fmt.Printf("Model: %s\n", res.Data.Name)
		fmt.Printf("ID: %s\n", res.Data.ID)
	}
}

// Example_streamChat demonstrates streaming chat responses.
func Example_streamChat() {
	ctx := context.Background()
	sdk := openrouter.New(
		openrouter.WithSecurity(os.Getenv("OPENROUTER_API_KEY")),
	)

	res, err := sdk.Chat.Send(ctx, components.ChatRequest{
		Model: openrouter.Pointer("openai/gpt-4o"),
		Messages: []components.ChatMessages{
			components.CreateChatMessagesUser(
				components.ChatUserMessage{
					Role:    components.ChatUserMessageRoleUser,
					Content: components.CreateChatUserMessageContentStr("Explain quantum computing in simple terms"),
				},
			),
		},
		Stream: openrouter.Pointer(true),
	}, nil)
	if err != nil {
		log.Fatal(err)
	}

	if res != nil {
		// Handle streaming response
		fmt.Println("Streaming response received...")
	}
}

// Example_listProviders demonstrates listing available providers.
func Example_listProviders() {
	ctx := context.Background()
	sdk := openrouter.New(
		openrouter.WithSecurity(os.Getenv("OPENROUTER_API_KEY")),
	)

	res, err := sdk.Providers.List(ctx, nil)
	if err != nil {
		log.Fatal(err)
	}

	if res != nil && len(res.Data) > 0 {
		fmt.Printf("Found %d providers\n", len(res.Data))
		for i, provider := range res.Data {
			if i < 3 {
				fmt.Printf("- %s\n", provider.Name)
			}
		}
	}
}
