# Generations

## Overview

Generation history endpoints

### Available Operations

* [GetGeneration](#getgeneration) - Get request & usage metadata for a generation
* [ListGenerationContent](#listgenerationcontent) - Get stored prompt and completion content for a generation

## GetGeneration

Get request & usage metadata for a generation

### Example Usage

<!-- UsageSnippet language="go" operationID="getGeneration" method="get" path="/generation" -->
```go
package main

import(
	"context"
	"os"
	openrouter "github.com/OpenRouterTeam/go-sdk"
	"log"
)

func main() {
    ctx := context.Background()

    s := openrouter.New(
        openrouter.WithSecurity(os.Getenv("OPENROUTER_API_KEY")),
    )

    res, err := s.Generations.GetGeneration(ctx, "<id>")
    if err != nil {
        log.Fatal(err)
    }
    if res != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                | Type                                                     | Required                                                 | Description                                              | Example                                                  |
| -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- |
| `ctx`                                                    | [context.Context](https://pkg.go.dev/context#Context)    | :heavy_check_mark:                                       | The context to use for the request.                      |                                                          |
| `id`                                                     | `string`                                                 | :heavy_check_mark:                                       | The generation ID                                        | gen-1234567890                                           |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |                                                          |

### Response

**[*components.GenerationResponse](../../models/components/generationresponse.md), error**

### Errors

| Error Type                                | Status Code                               | Content Type                              |
| ----------------------------------------- | ----------------------------------------- | ----------------------------------------- |
| sdkerrors.UnauthorizedResponseError       | 401                                       | application/json                          |
| sdkerrors.PaymentRequiredResponseError    | 402                                       | application/json                          |
| sdkerrors.NotFoundResponseError           | 404                                       | application/json                          |
| sdkerrors.TooManyRequestsResponseError    | 429                                       | application/json                          |
| sdkerrors.InternalServerResponseError     | 500                                       | application/json                          |
| sdkerrors.BadGatewayResponseError         | 502                                       | application/json                          |
| sdkerrors.EdgeNetworkTimeoutResponseError | 524                                       | application/json                          |
| sdkerrors.ProviderOverloadedResponseError | 529                                       | application/json                          |
| sdkerrors.APIError                        | 4XX, 5XX                                  | \*/\*                                     |

## ListGenerationContent

Get stored prompt and completion content for a generation

### Example Usage

<!-- UsageSnippet language="go" operationID="listGenerationContent" method="get" path="/generation/content" -->
```go
package main

import(
	"context"
	"os"
	openrouter "github.com/OpenRouterTeam/go-sdk"
	"log"
)

func main() {
    ctx := context.Background()

    s := openrouter.New(
        openrouter.WithSecurity(os.Getenv("OPENROUTER_API_KEY")),
    )

    res, err := s.Generations.ListGenerationContent(ctx, "gen-1234567890")
    if err != nil {
        log.Fatal(err)
    }
    if res != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                | Type                                                     | Required                                                 | Description                                              | Example                                                  |
| -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- |
| `ctx`                                                    | [context.Context](https://pkg.go.dev/context#Context)    | :heavy_check_mark:                                       | The context to use for the request.                      |                                                          |
| `id`                                                     | `string`                                                 | :heavy_check_mark:                                       | The generation ID                                        | gen-1234567890                                           |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |                                                          |

### Response

**[*components.GenerationContentResponse](../../models/components/generationcontentresponse.md), error**

### Errors

| Error Type                                | Status Code                               | Content Type                              |
| ----------------------------------------- | ----------------------------------------- | ----------------------------------------- |
| sdkerrors.UnauthorizedResponseError       | 401                                       | application/json                          |
| sdkerrors.ForbiddenResponseError          | 403                                       | application/json                          |
| sdkerrors.NotFoundResponseError           | 404                                       | application/json                          |
| sdkerrors.TooManyRequestsResponseError    | 429                                       | application/json                          |
| sdkerrors.InternalServerResponseError     | 500                                       | application/json                          |
| sdkerrors.BadGatewayResponseError         | 502                                       | application/json                          |
| sdkerrors.EdgeNetworkTimeoutResponseError | 524                                       | application/json                          |
| sdkerrors.ProviderOverloadedResponseError | 529                                       | application/json                          |
| sdkerrors.APIError                        | 4XX, 5XX                                  | \*/\*                                     |