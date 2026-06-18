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

	res, err := s.Providers.List(ctx)
	if err != nil {
		log.Fatal(err)
	}
	if res == nil {
		log.Fatal("empty providers response")
	}

	fmt.Printf("%d providers (showing first %d)\n", len(res.Data), listLimit)
	for i, provider := range res.Data {
		if i >= listLimit {
			break
		}
		fmt.Printf("  %-20s  %s\n", provider.Slug, provider.Name)
	}
}
