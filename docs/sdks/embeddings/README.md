# Embeddings

## Overview

Text embedding endpoints

### Available Operations

* [Generate](#generate) - Submit an embedding request
* [ListModels](#listmodels) - List all embeddings models

## Generate

Submits an embedding request to the embeddings router

### Example Usage

<!-- UsageSnippet language="go" operationID="createEmbeddings" method="post" path="/embeddings" -->
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

    res, err := s.Embeddings.Generate(ctx, operations.CreateEmbeddingsRequest{
        Input: operations.CreateInputUnionStr(
            "<value>",
        ),
        Model: "Taurus",
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

| Parameter                                                                                | Type                                                                                     | Required                                                                                 | Description                                                                              |
| ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- |
| `ctx`                                                                                    | [context.Context](https://pkg.go.dev/context#Context)                                    | :heavy_check_mark:                                                                       | The context to use for the request.                                                      |
| `request`                                                                                | [operations.CreateEmbeddingsRequest](../../models/operations/createembeddingsrequest.md) | :heavy_check_mark:                                                                       | The request object to use for the request.                                               |
| `opts`                                                                                   | [][operations.Option](../../models/operations/option.md)                                 | :heavy_minus_sign:                                                                       | The options for this request.                                                            |

### Response

**[*operations.CreateEmbeddingsResponse](../../models/operations/createembeddingsresponse.md), error**

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

## ListModels

Returns a list of all available embeddings models and their properties

### Example Usage

<!-- UsageSnippet language="go" operationID="listEmbeddingsModels" method="get" path="/embeddings/models" -->
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

    res, err := s.Embeddings.ListModels(ctx)
    if err != nil {
        log.Fatal(err)
    }
    if res != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                | Type                                                     | Required                                                 | Description                                              |
| -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- |
| `ctx`                                                    | [context.Context](https://pkg.go.dev/context#Context)    | :heavy_check_mark:                                       | The context to use for the request.                      |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |

### Response

**[*components.ModelsListResponse](../../models/components/modelslistresponse.md), error**

### Errors

| Error Type                            | Status Code                           | Content Type                          |
| ------------------------------------- | ------------------------------------- | ------------------------------------- |
| sdkerrors.BadRequestResponseError     | 400                                   | application/json                      |
| sdkerrors.InternalServerResponseError | 500                                   | application/json                      |
| sdkerrors.APIError                    | 4XX, 5XX                              | \*/\*                                 |