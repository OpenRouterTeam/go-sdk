# Beta.Responses

## Overview

beta.responses endpoints

### Available Operations

* [Send](#send) - Create a response

## Send

Creates a streaming or non-streaming response using OpenResponses API format

### Example Usage: guardrail-blocked

<!-- UsageSnippet language="go" operationID="createResponses" method="post" path="/responses" example="guardrail-blocked" -->
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

    res, err := s.Beta.Responses.Send(ctx, components.ResponsesRequest{}, nil)
    if err != nil {
        log.Fatal(err)
    }
    if res != nil {
        defer res.ResponsesStreamingResponse.Close()

        for res.ResponsesStreamingResponse.Next() {
            event := res.ResponsesStreamingResponse.Value()
            log.Print(event)
            // Handle the event
	      }
    }
}
```
### Example Usage: insufficient-permissions

<!-- UsageSnippet language="go" operationID="createResponses" method="post" path="/responses" example="insufficient-permissions" -->
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

    res, err := s.Beta.Responses.Send(ctx, components.ResponsesRequest{}, nil)
    if err != nil {
        log.Fatal(err)
    }
    if res != nil {
        defer res.ResponsesStreamingResponse.Close()

        for res.ResponsesStreamingResponse.Next() {
            event := res.ResponsesStreamingResponse.Value()
            log.Print(event)
            // Handle the event
	      }
    }
}
```

### Parameters

| Parameter                                                                                                                                                                                                                                                                                                                                                                                             | Type                                                                                                                                                                                                                                                                                                                                                                                                  | Required                                                                                                                                                                                                                                                                                                                                                                                              | Description                                                                                                                                                                                                                                                                                                                                                                                           | Example                                                                                                                                                                                                                                                                                                                                                                                               |
| ----------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | ----------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | ----------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | ----------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | ----------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                                                                                                                                                                                                                                                                                                                 | [context.Context](https://pkg.go.dev/context#Context)                                                                                                                                                                                                                                                                                                                                                 | :heavy_check_mark:                                                                                                                                                                                                                                                                                                                                                                                    | The context to use for the request.                                                                                                                                                                                                                                                                                                                                                                   |                                                                                                                                                                                                                                                                                                                                                                                                       |
| `responsesRequest`                                                                                                                                                                                                                                                                                                                                                                                    | [components.ResponsesRequest](../../models/components/responsesrequest.md)                                                                                                                                                                                                                                                                                                                            | :heavy_check_mark:                                                                                                                                                                                                                                                                                                                                                                                    | N/A                                                                                                                                                                                                                                                                                                                                                                                                   | {<br/>"input": [<br/>{<br/>"content": "Hello, how are you?",<br/>"role": "user",<br/>"type": "message"<br/>}<br/>],<br/>"model": "anthropic/claude-4.5-sonnet-20250929",<br/>"temperature": 0.7,<br/>"tools": [<br/>{<br/>"description": "Get the current weather in a given location",<br/>"name": "get_current_weather",<br/>"parameters": {<br/>"properties": {<br/>"location": {<br/>"type": "string"<br/>}<br/>},<br/>"type": "object"<br/>},<br/>"type": "function"<br/>}<br/>],<br/>"top_p": 0.9<br/>} |
| `xOpenRouterMetadata`                                                                                                                                                                                                                                                                                                                                                                                 | [*components.MetadataLevel](../../models/components/metadatalevel.md)                                                                                                                                                                                                                                                                                                                                 | :heavy_minus_sign:                                                                                                                                                                                                                                                                                                                                                                                    | Opt-in to surface routing metadata on the response under `openrouter_metadata`. Defaults to `disabled`. The legacy header `X-OpenRouter-Experimental-Metadata` is also accepted for backward compatibility.                                                                                                                                                                                           | enabled                                                                                                                                                                                                                                                                                                                                                                                               |
| `opts`                                                                                                                                                                                                                                                                                                                                                                                                | [][operations.Option](../../models/operations/option.md)                                                                                                                                                                                                                                                                                                                                              | :heavy_minus_sign:                                                                                                                                                                                                                                                                                                                                                                                    | The options for this request.                                                                                                                                                                                                                                                                                                                                                                         |                                                                                                                                                                                                                                                                                                                                                                                                       |

### Response

**[*operations.CreateResponsesResponse](../../models/operations/createresponsesresponse.md), error**

### Errors

| Error Type                                 | Status Code                                | Content Type                               |
| ------------------------------------------ | ------------------------------------------ | ------------------------------------------ |
| sdkerrors.BadRequestResponseError          | 400                                        | application/json                           |
| sdkerrors.UnauthorizedResponseError        | 401                                        | application/json                           |
| sdkerrors.PaymentRequiredResponseError     | 402                                        | application/json                           |
| sdkerrors.ForbiddenResponseError           | 403                                        | application/json                           |
| sdkerrors.NotFoundResponseError            | 404                                        | application/json                           |
| sdkerrors.RequestTimeoutResponseError      | 408                                        | application/json                           |
| sdkerrors.PayloadTooLargeResponseError     | 413                                        | application/json                           |
| sdkerrors.UnprocessableEntityResponseError | 422                                        | application/json                           |
| sdkerrors.TooManyRequestsResponseError     | 429                                        | application/json                           |
| sdkerrors.InternalServerResponseError      | 500                                        | application/json                           |
| sdkerrors.BadGatewayResponseError          | 502                                        | application/json                           |
| sdkerrors.ServiceUnavailableResponseError  | 503                                        | application/json                           |
| sdkerrors.EdgeNetworkTimeoutResponseError  | 524                                        | application/json                           |
| sdkerrors.ProviderOverloadedResponseError  | 529                                        | application/json                           |
| sdkerrors.APIError                         | 4XX, 5XX                                   | \*/\*                                      |