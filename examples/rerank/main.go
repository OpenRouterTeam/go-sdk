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

	res, err := s.Rerank.Rerank(ctx, operations.CreateRerankRequest{
		Model: "cohere/rerank-v3.5",
		Query: "capital of France",
		Documents: []operations.Document{
			operations.CreateDocumentStr("Paris is the capital of France."),
			operations.CreateDocumentStr("Berlin is the capital of Germany."),
			operations.CreateDocumentStr("Madrid is the capital of Spain."),
		},
		TopN: openrouter.Pointer(int64(3)),
	})
	if err != nil {
		log.Fatal(err)
	}
	if res == nil || res.CreateRerankResponseBody == nil {
		log.Fatal("empty rerank response")
	}

	body := res.CreateRerankResponseBody
	fmt.Printf("model: %s\n", body.Model)
	for _, result := range body.Results {
		text := ""
		if result.Document.Text != nil {
			text = *result.Document.Text
		}
		fmt.Printf("  #%d score=%.4f  %s\n", result.Index, result.RelevanceScore, text)
	}
}
