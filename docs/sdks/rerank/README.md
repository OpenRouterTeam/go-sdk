# Rerank

## Overview

Rerank endpoints

### Available Operations

* [Rerank](#rerank) - Submit a rerank request

## Rerank

Submits a rerank request to the rerank router

### Example Usage

<!-- UsageSnippet language="go" operationID="createRerank" method="post" path="/rerank" -->
```go
package main

import(
	"context"
	"os"
	openrouter "github.com/OpenRouterTeam/go-sdk"
	"github.com/OpenRouterTeam/go-sdk/models/operations"
	"log"
)

func main() {
    ctx := context.Background()

    s := openrouter.New(
        openrouter.WithSecurity(os.Getenv("OPENROUTER_API_KEY")),
    )

    res, err := s.Rerank.Rerank(ctx, operations.CreateRerankRequest{
        Documents: []string{
            "Paris is the capital of France.",
            "Berlin is the capital of Germany.",
        },
        Model: "cohere/rerank-v3.5",
        Query: "What is the capital of France?",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                        | Type                                                                             | Required                                                                         | Description                                                                      |
| -------------------------------------------------------------------------------- | -------------------------------------------------------------------------------- | -------------------------------------------------------------------------------- | -------------------------------------------------------------------------------- |
| `ctx`                                                                            | [context.Context](https://pkg.go.dev/context#Context)                            | :heavy_check_mark:                                                               | The context to use for the request.                                              |
| `request`                                                                        | [operations.CreateRerankRequest](../../models/operations/creatererankrequest.md) | :heavy_check_mark:                                                               | The request object to use for the request.                                       |
| `opts`                                                                           | [][operations.Option](../../models/operations/option.md)                         | :heavy_minus_sign:                                                               | The options for this request.                                                    |

### Response

**[*operations.CreateRerankResponse](../../models/operations/creatererankresponse.md), error**

### Errors

| Error Type                                | Status Code                               | Content Type                              |
| ----------------------------------------- | ----------------------------------------- | ----------------------------------------- |
| sdkerrors.BadRequestResponseError         | 400                                       | application/json                          |
| sdkerrors.UnauthorizedResponseError       | 401                                       | application/json                          |
| sdkerrors.PaymentRequiredResponseError    | 402                                       | application/json                          |
| sdkerrors.NotFoundResponseError           | 404                                       | application/json                          |
| sdkerrors.TooManyRequestsResponseError    | 429                                       | application/json                          |
| sdkerrors.InternalServerResponseError     | 500                                       | application/json                          |
| sdkerrors.BadGatewayResponseError         | 502                                       | application/json                          |
| sdkerrors.ServiceUnavailableResponseError | 503                                       | application/json                          |
| sdkerrors.EdgeNetworkTimeoutResponseError | 524                                       | application/json                          |
| sdkerrors.ProviderOverloadedResponseError | 529                                       | application/json                          |
| sdkerrors.APIError                        | 4XX, 5XX                                  | \*/\*                                     |