# Classifications

## Overview

Task classification market-share endpoints

### Available Operations

* [GetTaskClassifications](#gettaskclassifications) - Task classification market share

## GetTaskClassifications

Returns the market-share breakdown of OpenRouter traffic by task classification
(e.g. code generation, web search, summarization) over a trailing time window.

Each classification reports its share of classified sampled requests (`usage_share`)
and classified sampled token volume (`token_share`) as fractions between 0 and 1.
The unclassified `other` bucket is excluded. Absolute volumes are not exposed
because the underlying data is sampled.

Each classification also includes a `models` array listing the top models by
request volume within that classification, with their within-tag usage and token shares.

Classifications are grouped into macro-categories (Code, Data, Agent, General)
with aggregate shares provided for each.

Authenticate with any valid OpenRouter API key (same key used for inference).
Rate-limited to 30 requests/minute per key and 500 requests/day per account.

When republishing or quoting this data, cite as:
"Source: OpenRouter (openrouter.ai/rankings), as of {as_of}."

### Example Usage

<!-- UsageSnippet language="go" operationID="getTaskClassifications" method="get" path="/classifications/task" -->
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

    res, err := s.Classifications.GetTaskClassifications(ctx, operations.WindowSevend.ToPointer())
    if err != nil {
        log.Fatal(err)
    }
    if res != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                             | Type                                                                                                  | Required                                                                                              | Description                                                                                           | Example                                                                                               |
| ----------------------------------------------------------------------------------------------------- | ----------------------------------------------------------------------------------------------------- | ----------------------------------------------------------------------------------------------------- | ----------------------------------------------------------------------------------------------------- | ----------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                 | [context.Context](https://pkg.go.dev/context#Context)                                                 | :heavy_check_mark:                                                                                    | The context to use for the request.                                                                   |                                                                                                       |
| `window`                                                                                              | [*operations.Window](../../models/operations/window.md)                                               | :heavy_minus_sign:                                                                                    | Trailing time window for the classification data. Currently only `7d` (trailing 7 days) is supported. | 7d                                                                                                    |
| `opts`                                                                                                | [][operations.Option](../../models/operations/option.md)                                              | :heavy_minus_sign:                                                                                    | The options for this request.                                                                         |                                                                                                       |

### Response

**[*components.TaskClassificationResponse](../../models/components/taskclassificationresponse.md), error**

### Errors

| Error Type                             | Status Code                            | Content Type                           |
| -------------------------------------- | -------------------------------------- | -------------------------------------- |
| sdkerrors.BadRequestResponseError      | 400                                    | application/json                       |
| sdkerrors.UnauthorizedResponseError    | 401                                    | application/json                       |
| sdkerrors.TooManyRequestsResponseError | 429                                    | application/json                       |
| sdkerrors.InternalServerResponseError  | 500                                    | application/json                       |
| sdkerrors.APIError                     | 4XX, 5XX                               | \*/\*                                  |