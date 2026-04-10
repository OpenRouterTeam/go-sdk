# VideoGeneration

## Overview

Video Generation endpoints

### Available Operations

* [Generate](#generate) - Submit a video generation request
* [GetGeneration](#getgeneration) - Poll video generation status
* [GetVideoContent](#getvideocontent) - Download generated video content
* [ListVideosModels](#listvideosmodels) - List all video generation models

## Generate

Submits a video generation request and returns a polling URL to check status

### Example Usage

<!-- UsageSnippet language="go" operationID="createVideos" method="post" path="/videos" -->
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

    res, err := s.VideoGeneration.Generate(ctx, components.VideoGenerationRequest{
        AspectRatio: components.AspectRatioOneHundredAndSixtyNine.ToPointer(),
        Duration: openrouter.Pointer[int64](8),
        Model: "google/veo-3.1",
        Prompt: "A serene mountain landscape at sunset",
        Resolution: components.ResolutionSevenHundredAndTwentyp.ToPointer(),
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
| `request`                                                                              | [components.VideoGenerationRequest](../../models/components/videogenerationrequest.md) | :heavy_check_mark:                                                                     | The request object to use for the request.                                             |
| `opts`                                                                                 | [][operations.Option](../../models/operations/option.md)                               | :heavy_minus_sign:                                                                     | The options for this request.                                                          |

### Response

**[*components.VideoGenerationResponse](../../models/components/videogenerationresponse.md), error**

### Errors

| Error Type                             | Status Code                            | Content Type                           |
| -------------------------------------- | -------------------------------------- | -------------------------------------- |
| sdkerrors.BadRequestResponseError      | 400                                    | application/json                       |
| sdkerrors.UnauthorizedResponseError    | 401                                    | application/json                       |
| sdkerrors.PaymentRequiredResponseError | 402                                    | application/json                       |
| sdkerrors.NotFoundResponseError        | 404                                    | application/json                       |
| sdkerrors.TooManyRequestsResponseError | 429                                    | application/json                       |
| sdkerrors.InternalServerResponseError  | 500                                    | application/json                       |
| sdkerrors.APIError                     | 4XX, 5XX                               | \*/\*                                  |

## GetGeneration

Returns job status and content URLs when completed

### Example Usage

<!-- UsageSnippet language="go" operationID="getVideos" method="get" path="/videos/{jobId}" -->
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

    res, err := s.VideoGeneration.GetGeneration(ctx, "job-abc123")
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
| `jobID`                                                  | `string`                                                 | :heavy_check_mark:                                       | N/A                                                      | job-abc123                                               |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |                                                          |

### Response

**[*components.VideoGenerationResponse](../../models/components/videogenerationresponse.md), error**

### Errors

| Error Type                            | Status Code                           | Content Type                          |
| ------------------------------------- | ------------------------------------- | ------------------------------------- |
| sdkerrors.UnauthorizedResponseError   | 401                                   | application/json                      |
| sdkerrors.NotFoundResponseError       | 404                                   | application/json                      |
| sdkerrors.InternalServerResponseError | 500                                   | application/json                      |
| sdkerrors.APIError                    | 4XX, 5XX                              | \*/\*                                 |

## GetVideoContent

Streams the generated video content from the upstream provider

### Example Usage

<!-- UsageSnippet language="go" operationID="listVideosContent" method="get" path="/videos/{jobId}/content" -->
```go
package main

import(
	"context"
	"os"
	openrouter "github.com/OpenRouterTeam/go-sdk"
	"github.com/OpenRouterTeam/go-sdk/optionalnullable"
	"log"
)

func main() {
    ctx := context.Background()

    s := openrouter.New(
        openrouter.WithSecurity(os.Getenv("OPENROUTER_API_KEY")),
    )

    res, err := s.VideoGeneration.GetVideoContent(ctx, "job-abc123", optionalnullable.From(openrouter.Pointer[int64](0)))
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
| `jobID`                                                  | `string`                                                 | :heavy_check_mark:                                       | N/A                                                      | job-abc123                                               |
| `index`                                                  | optionalnullable.OptionalNullable[`int64`]               | :heavy_minus_sign:                                       | N/A                                                      | 0                                                        |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |                                                          |

### Response

**[io.ReadCloser](../../.md), error**

### Errors

| Error Type                            | Status Code                           | Content Type                          |
| ------------------------------------- | ------------------------------------- | ------------------------------------- |
| sdkerrors.BadRequestResponseError     | 400                                   | application/json                      |
| sdkerrors.UnauthorizedResponseError   | 401                                   | application/json                      |
| sdkerrors.NotFoundResponseError       | 404                                   | application/json                      |
| sdkerrors.InternalServerResponseError | 500                                   | application/json                      |
| sdkerrors.BadGatewayResponseError     | 502                                   | application/json                      |
| sdkerrors.APIError                    | 4XX, 5XX                              | \*/\*                                 |

## ListVideosModels

Returns a list of all available video generation models and their properties

### Example Usage

<!-- UsageSnippet language="go" operationID="listVideosModels" method="get" path="/videos/models" -->
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

    res, err := s.VideoGeneration.ListVideosModels(ctx)
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

**[*components.VideoModelsListResponse](../../models/components/videomodelslistresponse.md), error**

### Errors

| Error Type                            | Status Code                           | Content Type                          |
| ------------------------------------- | ------------------------------------- | ------------------------------------- |
| sdkerrors.BadRequestResponseError     | 400                                   | application/json                      |
| sdkerrors.InternalServerResponseError | 500                                   | application/json                      |
| sdkerrors.APIError                    | 4XX, 5XX                              | \*/\*                                 |