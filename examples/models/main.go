package main

import (
	"context"
	"fmt"
	"log"
	"os"

	openrouter "github.com/OpenRouterTeam/go-sdk"
)

const listLimit = 10

func main() {
	ctx := context.Background()

	s := openrouter.New(
		openrouter.WithSecurity(os.Getenv("OPENROUTER_API_KEY")),
	)

	res, err := s.Models.List(ctx, nil)
	if err != nil {
		log.Fatal(err)
	}
	if res == nil {
		log.Fatal("empty models response")
	}

	fmt.Printf("%d models (showing first %d)\n", len(res.Result.Data), listLimit)
	for i, model := range res.Result.Data {
		if i >= listLimit {
			break
		}
		contextLen := "?"
		if model.ContextLength != nil {
			contextLen = fmt.Sprintf("%d", *model.ContextLength)
		}
		fmt.Printf("  %-40s  ctx=%-8s  %s\n", model.ID, contextLen, model.Name)
	}
}
