# Chat

## Overview

### Available Operations

* [Send](#send) - Create a chat completion

## Send

Sends a request for a model response for the given chat conversation. Supports both streaming and non-streaming modes.

### Example Usage

<!-- UsageSnippet language="go" operationID="sendChatCompletionRequest" method="post" path="/chat/completions" -->
```go
package main

import(
	"context"
	"os"
	openrouter "github.com/OpenRouterTeam/go-sdk"
	"github.com/OpenRouterTeam/go-sdk/models/components"
	"log"
)

func main() {
    ctx := context.Background()

    s := openrouter.New(
        openrouter.WithSecurity(os.Getenv("OPENROUTER_API_KEY")),
    )

    res, err := s.Chat.Send(ctx, components.ChatGenerationParams{
        Messages: []components.Message{},
    })
    if err != nil {
        log.Fatal(err)
    }
    if res != nil {
        defer res.ChatStreamingResponseChunk.Close()

        for res.ChatStreamingResponseChunk.Next() {
            event := res.ChatStreamingResponseChunk.Value()
            log.Print(event)
            // Handle the event
	      }
    }
}
```

### Parameters

| Parameter                                                                          | Type                                                                               | Required                                                                           | Description                                                                        |
| ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- |
| `ctx`                                                                              | [context.Context](https://pkg.go.dev/context#Context)                              | :heavy_check_mark:                                                                 | The context to use for the request.                                                |
| `request`                                                                          | [components.ChatGenerationParams](../../models/components/chatgenerationparams.md) | :heavy_check_mark:                                                                 | The request object to use for the request.                                         |
| `opts`                                                                             | [][operations.Option](../../models/operations/option.md)                           | :heavy_minus_sign:                                                                 | The options for this request.                                                      |

### Response

**[*operations.SendChatCompletionRequestResponse](../../models/operations/sendchatcompletionrequestresponse.md), error**

### Errors

| Error Type          | Status Code         | Content Type        |
| ------------------- | ------------------- | ------------------- |
| sdkerrors.ChatError | 400, 401, 429       | application/json    |
| sdkerrors.ChatError | 500                 | application/json    |
| sdkerrors.APIError  | 4XX, 5XX            | \*/\*               |