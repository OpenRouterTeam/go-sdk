# Models

## Overview

Model information endpoints

### Available Operations

* [List](#list) - List all models and their properties
* [Count](#count) - Get total count of available models
* [ListForUser](#listforuser) - List models filtered by user provider preferences, privacy settings, and guardrails

## List

List all models and their properties

### Example Usage

<!-- UsageSnippet language="go" operationID="getModels" method="get" path="/models" -->
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

    res, err := s.Models.List(ctx, nil, nil, nil)
    if err != nil {
        log.Fatal(err)
    }
    if res != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                                                                           | Type                                                                                                                                                                | Required                                                                                                                                                            | Description                                                                                                                                                         | Example                                                                                                                                                             |
| ------------------------------------------------------------------------------------------------------------------------------------------------------------------- | ------------------------------------------------------------------------------------------------------------------------------------------------------------------- | ------------------------------------------------------------------------------------------------------------------------------------------------------------------- | ------------------------------------------------------------------------------------------------------------------------------------------------------------------- | ------------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                                                                               | [context.Context](https://pkg.go.dev/context#Context)                                                                                                               | :heavy_check_mark:                                                                                                                                                  | The context to use for the request.                                                                                                                                 |                                                                                                                                                                     |
| `category`                                                                                                                                                          | [*operations.Category](../../models/operations/category.md)                                                                                                         | :heavy_minus_sign:                                                                                                                                                  | Filter models by use case category                                                                                                                                  | programming                                                                                                                                                         |
| `supportedParameters`                                                                                                                                               | `*string`                                                                                                                                                           | :heavy_minus_sign:                                                                                                                                                  | Filter models by supported parameter (comma-separated)                                                                                                              | temperature                                                                                                                                                         |
| `outputModalities`                                                                                                                                                  | `*string`                                                                                                                                                           | :heavy_minus_sign:                                                                                                                                                  | Filter models by output modality. Accepts a comma-separated list of modalities (text, image, audio, embeddings) or "all" to include all models. Defaults to "text". | text                                                                                                                                                                |
| `opts`                                                                                                                                                              | [][operations.Option](../../models/operations/option.md)                                                                                                            | :heavy_minus_sign:                                                                                                                                                  | The options for this request.                                                                                                                                       |                                                                                                                                                                     |

### Response

**[*components.ModelsListResponse](../../models/components/modelslistresponse.md), error**

### Errors

| Error Type                            | Status Code                           | Content Type                          |
| ------------------------------------- | ------------------------------------- | ------------------------------------- |
| sdkerrors.BadRequestResponseError     | 400                                   | application/json                      |
| sdkerrors.InternalServerResponseError | 500                                   | application/json                      |
| sdkerrors.APIError                    | 4XX, 5XX                              | \*/\*                                 |

## Count

Get total count of available models

### Example Usage

<!-- UsageSnippet language="go" operationID="listModelsCount" method="get" path="/models/count" -->
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

    res, err := s.Models.Count(ctx, nil)
    if err != nil {
        log.Fatal(err)
    }
    if res != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                                                                           | Type                                                                                                                                                                | Required                                                                                                                                                            | Description                                                                                                                                                         | Example                                                                                                                                                             |
| ------------------------------------------------------------------------------------------------------------------------------------------------------------------- | ------------------------------------------------------------------------------------------------------------------------------------------------------------------- | ------------------------------------------------------------------------------------------------------------------------------------------------------------------- | ------------------------------------------------------------------------------------------------------------------------------------------------------------------- | ------------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                                                                               | [context.Context](https://pkg.go.dev/context#Context)                                                                                                               | :heavy_check_mark:                                                                                                                                                  | The context to use for the request.                                                                                                                                 |                                                                                                                                                                     |
| `outputModalities`                                                                                                                                                  | `*string`                                                                                                                                                           | :heavy_minus_sign:                                                                                                                                                  | Filter models by output modality. Accepts a comma-separated list of modalities (text, image, audio, embeddings) or "all" to include all models. Defaults to "text". | text                                                                                                                                                                |
| `opts`                                                                                                                                                              | [][operations.Option](../../models/operations/option.md)                                                                                                            | :heavy_minus_sign:                                                                                                                                                  | The options for this request.                                                                                                                                       |                                                                                                                                                                     |

### Response

**[*components.ModelsCountResponse](../../models/components/modelscountresponse.md), error**

### Errors

| Error Type                            | Status Code                           | Content Type                          |
| ------------------------------------- | ------------------------------------- | ------------------------------------- |
| sdkerrors.BadRequestResponseError     | 400                                   | application/json                      |
| sdkerrors.InternalServerResponseError | 500                                   | application/json                      |
| sdkerrors.APIError                    | 4XX, 5XX                              | \*/\*                                 |

## ListForUser

List models filtered by user provider preferences, [privacy settings](https://openrouter.ai/docs/guides/privacy/provider-logging), and [guardrails](https://openrouter.ai/docs/guides/features/guardrails). If requesting through `eu.openrouter.ai/api/v1/...` the results will be filtered to models that satisfy [EU in-region routing](https://openrouter.ai/docs/guides/privacy/provider-logging#enterprise-eu-in-region-routing).

### Example Usage

<!-- UsageSnippet language="go" operationID="listModelsUser" method="get" path="/models/user" -->
```go
package main

import(
	"context"
	openrouter "github.com/OpenRouterTeam/go-sdk"
	"os"
	"github.com/OpenRouterTeam/go-sdk/models/operations"
	"log"
)

func main() {
    ctx := context.Background()

    s := openrouter.New()

    res, err := s.Models.ListForUser(ctx, operations.ListModelsUserSecurity{
        Bearer: os.Getenv("OPENROUTER_BEARER"),
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

| Parameter                                                                              | Type                                                                                   | Required                                                                               | Description                                                                            |
| -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- |
| `ctx`                                                                                  | [context.Context](https://pkg.go.dev/context#Context)                                  | :heavy_check_mark:                                                                     | The context to use for the request.                                                    |
| `security`                                                                             | [operations.ListModelsUserSecurity](../../models/operations/listmodelsusersecurity.md) | :heavy_check_mark:                                                                     | The security requirements to use for the request.                                      |
| `opts`                                                                                 | [][operations.Option](../../models/operations/option.md)                               | :heavy_minus_sign:                                                                     | The options for this request.                                                          |

### Response

**[*components.ModelsListResponse](../../models/components/modelslistresponse.md), error**

### Errors

| Error Type                            | Status Code                           | Content Type                          |
| ------------------------------------- | ------------------------------------- | ------------------------------------- |
| sdkerrors.UnauthorizedResponseError   | 401                                   | application/json                      |
| sdkerrors.NotFoundResponseError       | 404                                   | application/json                      |
| sdkerrors.InternalServerResponseError | 500                                   | application/json                      |
| sdkerrors.APIError                    | 4XX, 5XX                              | \*/\*                                 |