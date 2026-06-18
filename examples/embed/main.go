package main

import (
	"context"
	"fmt"
	"log"
	"os"

	openrouter "github.com/OpenRouterTeam/go-sdk"
	"github.com/OpenRouterTeam/go-sdk/models/operations"
)

func main() {
	ctx := context.Background()

	s := openrouter.New(
		openrouter.WithSecurity(os.Getenv("OPENROUTER_API_KEY")),
	)

	res, err := s.Embeddings.Generate(ctx, operations.CreateEmbeddingsRequest{
		Model: "openai/text-embedding-3-small",
		Input: operations.CreateInputUnionStr("OpenRouter Go SDK embed example"),
	})
	if err != nil {
		log.Fatal(err)
	}
	if res == nil || res.CreateEmbeddingsResponseBody == nil {
		log.Fatal("empty embeddings response")
	}

	body := res.CreateEmbeddingsResponseBody
	fmt.Printf("model: %s\n", body.Model)
	if body.Usage != nil {
		fmt.Printf("tokens: %d\n", body.Usage.TotalTokens)
	}

	for _, item := range body.Data {
		if item.Embedding.Type == operations.EmbeddingTypeArrayOfNumber {
			fmt.Printf("embedding dimensions: %d\n", len(item.Embedding.ArrayOfNumber))
		}
	}
}
