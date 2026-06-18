package main

import (
	"context"
	"fmt"
	"log"
	"os"

	openrouter "github.com/OpenRouterTeam/go-sdk"
)

func main() {
	id := os.Getenv("OPENROUTER_GENERATION_ID")
	if id == "" {
		log.Fatal("OPENROUTER_GENERATION_ID is required (run examples/chat and use the response id)")
	}

	ctx := context.Background()

	s := openrouter.New(
		openrouter.WithSecurity(os.Getenv("OPENROUTER_API_KEY")),
	)

	res, err := s.Generations.GetGeneration(ctx, id)
	if err != nil {
		log.Fatal(err)
	}
	if res == nil {
		log.Fatal("empty generation response")
	}

	data := res.Data
	fmt.Printf("id: %s\n", data.ID)
	fmt.Printf("model: %s\n", data.Model)
	fmt.Printf("created_at: %s\n", data.CreatedAt)
	if data.ProviderName != nil {
		fmt.Printf("provider: %s\n", *data.ProviderName)
	}
	fmt.Printf("total_cost: $%.6f\n", data.TotalCost)
	if data.FinishReason != nil {
		fmt.Printf("finish_reason: %s\n", *data.FinishReason)
	}
}
