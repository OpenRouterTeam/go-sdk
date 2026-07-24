/*
Package openrouter provides an OpenAI-compatible API client with additional OpenRouter features.

The OpenRouter API offers OpenAI-compatible endpoints with additional features like model routing,
provider selection, and unified billing.

This SDK is in beta. Pin to a specific module version to avoid unexpected breaking changes:

	go get github.com/OpenRouterTeam/go-sdk@v0.5.39

For full API documentation, visit: https://openrouter.ai/docs/client-sdks/go/overview

Authentication:

	import (
	    "context"

	    openrouter "github.com/OpenRouterTeam/go-sdk"
	    "github.com/OpenRouterTeam/go-sdk/models/components"
	)

	sdk := openrouter.New(
	    openrouter.WithSecurity("your-api-key"),
	)

The API key can also be read from the OPENROUTER_API_KEY environment variable when using [New]
without [WithSecurity].

For license information, see the LICENSE file at the repository root.

Examples:

Basic chat completion:

	ctx := context.Background()
	res, err := sdk.Chat.Send(ctx, components.ChatRequest{
	    Model: openrouter.Pointer("openai/gpt-4o"),
	    Messages: []components.ChatMessages{
	        components.CreateChatMessagesUser(
	            components.ChatUserMessage{
	                Role: components.ChatUserMessageRoleUser,
	                Content: components.CreateChatUserMessageContentStr("Hello!"),
	            },
	        ),
	    },
	}, nil)
*/
package openrouter
