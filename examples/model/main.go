package main

import (
	"context"
	"fmt"
	"log"
	"os"

	openrouter "github.com/OpenRouterTeam/go-sdk"
)

func main() {
	ctx := context.Background()

	s := openrouter.New(
		openrouter.WithSecurity(os.Getenv("OPENROUTER_API_KEY")),
	)

	res, err := s.Models.Get(ctx, "openai", "gpt-4o-mini")
	if err != nil {
		log.Fatal(err)
	}
	if res == nil {
		log.Fatal("empty model response")
	}

	model := res.Data
	fmt.Printf("id: %s\n", model.ID)
	fmt.Printf("name: %s\n", model.Name)
	if model.Description != nil {
		fmt.Printf("description: %s\n", *model.Description)
	}
	if model.ContextLength != nil {
		fmt.Printf("context_length: %d\n", *model.ContextLength)
	}
	fmt.Printf("pricing: prompt=%s completion=%s\n",
		model.Pricing.Prompt,
		model.Pricing.Completion,
	)
}
