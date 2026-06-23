/*
Package openrouter provides an OpenAI-compatible API client with additional OpenRouter features.

The OpenRouter API offers OpenAI-compatible endpoints with additional features like model routing,
provider selection, and unified billing.

For full API documentation, visit: https://openrouter.ai/docs

Authentication:

	import openrouter "github.com/OpenRouterTeam/go-sdk"
	
	sdk := openrouter.New(
	    openrouter.WithSecurity("your-api-key"),
	)

For license information, see the LICENSE file at the repository root.

Examples:

Basic chat completion:

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
	})

See the example_test.go file for more comprehensive examples.
*/
package openrouter
