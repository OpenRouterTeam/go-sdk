# Byok

## Overview

BYOK endpoints

### Available Operations

* [List](#list) - List BYOK provider credentials
* [Create](#create) - Create a BYOK provider credential
* [Delete](#delete) - Delete a BYOK provider credential
* [Get](#get) - Get a BYOK provider credential
* [Update](#update) - Update a BYOK provider credential

## List

List the bring-your-own-key (BYOK) provider credentials for the authenticated entity's default workspace. Use the `workspace_id` query parameter to scope the result to a different workspace, or the `provider` query parameter to filter by upstream provider. [Management key](/docs/guides/overview/auth/management-api-keys) required.

### Example Usage

<!-- UsageSnippet language="go" operationID="listBYOKKeys" method="get" path="/byok" -->
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

    res, err := s.Byok.List(ctx, optionalnullable.From[int64](nil), nil, nil, nil)
    if err != nil {
        log.Fatal(err)
    }
    if res != nil {
        for {
            // handle items

            res, err = res.Next()

            if err != nil {
                // handle error
            }

            if res == nil {
                break
            }
        }
    }
}
```

### Parameters

| Parameter                                                                                     | Type                                                                                          | Required                                                                                      | Description                                                                                   | Example                                                                                       |
| --------------------------------------------------------------------------------------------- | --------------------------------------------------------------------------------------------- | --------------------------------------------------------------------------------------------- | --------------------------------------------------------------------------------------------- | --------------------------------------------------------------------------------------------- |
| `ctx`                                                                                         | [context.Context](https://pkg.go.dev/context#Context)                                         | :heavy_check_mark:                                                                            | The context to use for the request.                                                           |                                                                                               |
| `offset`                                                                                      | optionalnullable.OptionalNullable[`int64`]                                                    | :heavy_minus_sign:                                                                            | Number of records to skip for pagination                                                      | 0                                                                                             |
| `limit`                                                                                       | `*int64`                                                                                      | :heavy_minus_sign:                                                                            | Maximum number of records to return (max 100)                                                 | 50                                                                                            |
| `workspaceID`                                                                                 | `*string`                                                                                     | :heavy_minus_sign:                                                                            | Optional workspace ID to filter by. Defaults to the authenticated entity's default workspace. | 550e8400-e29b-41d4-a716-446655440000                                                          |
| `provider`                                                                                    | [*operations.Provider](../../models/operations/provider.md)                                   | :heavy_minus_sign:                                                                            | Optional provider slug to filter by (e.g. `openai`, `anthropic`, `amazon-bedrock`).           | openai                                                                                        |
| `opts`                                                                                        | [][operations.Option](../../models/operations/option.md)                                      | :heavy_minus_sign:                                                                            | The options for this request.                                                                 |                                                                                               |

### Response

**[*operations.ListBYOKKeysResponse](../../models/operations/listbyokkeysresponse.md), error**

### Errors

| Error Type                            | Status Code                           | Content Type                          |
| ------------------------------------- | ------------------------------------- | ------------------------------------- |
| sdkerrors.UnauthorizedResponseError   | 401                                   | application/json                      |
| sdkerrors.InternalServerResponseError | 500                                   | application/json                      |
| sdkerrors.APIError                    | 4XX, 5XX                              | \*/\*                                 |

## Create

Create a new bring-your-own-key (BYOK) provider credential. The raw key is encrypted at rest and never returned in API responses. Defaults to the authenticated entity's default workspace; use the `workspace_id` body field to scope to a different workspace. Treat the raw key as write-only; it is never returned after creation. [Management key](/docs/guides/overview/auth/management-api-keys) required.

### Example Usage

<!-- UsageSnippet language="go" operationID="createBYOKKey" method="post" path="/byok" -->
```go
package main

import(
	"context"
	"os"
	openrouter "github.com/OpenRouterTeam/go-sdk"
	"github.com/OpenRouterTeam/go-sdk/optionalnullable"
	"github.com/OpenRouterTeam/go-sdk/models/components"
	"log"
)

func main() {
    ctx := context.Background()

    s := openrouter.New(
        openrouter.WithSecurity(os.Getenv("OPENROUTER_API_KEY")),
    )

    res, err := s.Byok.Create(ctx, components.CreateBYOKKeyRequest{
        Key: "sk-proj-abc123...",
        Name: optionalnullable.From(openrouter.Pointer("Production OpenAI Key")),
        Provider: components.BYOKProviderSlugOpenai,
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

| Parameter                                                                          | Type                                                                               | Required                                                                           | Description                                                                        |
| ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- |
| `ctx`                                                                              | [context.Context](https://pkg.go.dev/context#Context)                              | :heavy_check_mark:                                                                 | The context to use for the request.                                                |
| `request`                                                                          | [components.CreateBYOKKeyRequest](../../models/components/createbyokkeyrequest.md) | :heavy_check_mark:                                                                 | The request object to use for the request.                                         |
| `opts`                                                                             | [][operations.Option](../../models/operations/option.md)                           | :heavy_minus_sign:                                                                 | The options for this request.                                                      |

### Response

**[*components.CreateBYOKKeyResponse](../../models/components/createbyokkeyresponse.md), error**

### Errors

| Error Type                            | Status Code                           | Content Type                          |
| ------------------------------------- | ------------------------------------- | ------------------------------------- |
| sdkerrors.BadRequestResponseError     | 400                                   | application/json                      |
| sdkerrors.UnauthorizedResponseError   | 401                                   | application/json                      |
| sdkerrors.ForbiddenResponseError      | 403                                   | application/json                      |
| sdkerrors.InternalServerResponseError | 500                                   | application/json                      |
| sdkerrors.APIError                    | 4XX, 5XX                              | \*/\*                                 |

## Delete

Delete (soft-delete) a bring-your-own-key (BYOK) provider credential by its `id`. The encrypted key material is wiped and the record is marked as deleted. [Management key](/docs/guides/overview/auth/management-api-keys) required.

### Example Usage

<!-- UsageSnippet language="go" operationID="deleteBYOKKey" method="delete" path="/byok/{id}" -->
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

    res, err := s.Byok.Delete(ctx, "11111111-2222-3333-4444-555555555555")
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
| `id`                                                     | `string`                                                 | :heavy_check_mark:                                       | The BYOK credential ID (UUID).                           | 11111111-2222-3333-4444-555555555555                     |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |                                                          |

### Response

**[*components.DeleteBYOKKeyResponse](../../models/components/deletebyokkeyresponse.md), error**

### Errors

| Error Type                            | Status Code                           | Content Type                          |
| ------------------------------------- | ------------------------------------- | ------------------------------------- |
| sdkerrors.UnauthorizedResponseError   | 401                                   | application/json                      |
| sdkerrors.NotFoundResponseError       | 404                                   | application/json                      |
| sdkerrors.InternalServerResponseError | 500                                   | application/json                      |
| sdkerrors.APIError                    | 4XX, 5XX                              | \*/\*                                 |

## Get

Get a single bring-your-own-key (BYOK) provider credential by its `id`. [Management key](/docs/guides/overview/auth/management-api-keys) required.

### Example Usage

<!-- UsageSnippet language="go" operationID="getBYOKKey" method="get" path="/byok/{id}" -->
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

    res, err := s.Byok.Get(ctx, "11111111-2222-3333-4444-555555555555")
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
| `id`                                                     | `string`                                                 | :heavy_check_mark:                                       | The BYOK credential ID (UUID).                           | 11111111-2222-3333-4444-555555555555                     |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |                                                          |

### Response

**[*components.GetBYOKKeyResponse](../../models/components/getbyokkeyresponse.md), error**

### Errors

| Error Type                            | Status Code                           | Content Type                          |
| ------------------------------------- | ------------------------------------- | ------------------------------------- |
| sdkerrors.UnauthorizedResponseError   | 401                                   | application/json                      |
| sdkerrors.NotFoundResponseError       | 404                                   | application/json                      |
| sdkerrors.InternalServerResponseError | 500                                   | application/json                      |
| sdkerrors.APIError                    | 4XX, 5XX                              | \*/\*                                 |

## Update

Update an existing bring-your-own-key (BYOK) provider credential by its `id`. Include the `key` field to rotate the raw provider API key in-place (the previous key material is overwritten). [Management key](/docs/guides/overview/auth/management-api-keys) required.

### Example Usage

<!-- UsageSnippet language="go" operationID="updateBYOKKey" method="patch" path="/byok/{id}" -->
```go
package main

import(
	"context"
	"os"
	openrouter "github.com/OpenRouterTeam/go-sdk"
	"github.com/OpenRouterTeam/go-sdk/optionalnullable"
	"github.com/OpenRouterTeam/go-sdk/models/components"
	"log"
)

func main() {
    ctx := context.Background()

    s := openrouter.New(
        openrouter.WithSecurity(os.Getenv("OPENROUTER_API_KEY")),
    )

    res, err := s.Byok.Update(ctx, "11111111-2222-3333-4444-555555555555", components.UpdateBYOKKeyRequest{
        Disabled: openrouter.Pointer(false),
        Name: optionalnullable.From(openrouter.Pointer("Updated OpenAI Key")),
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

| Parameter                                                                          | Type                                                                               | Required                                                                           | Description                                                                        | Example                                                                            |
| ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- |
| `ctx`                                                                              | [context.Context](https://pkg.go.dev/context#Context)                              | :heavy_check_mark:                                                                 | The context to use for the request.                                                |                                                                                    |
| `id`                                                                               | `string`                                                                           | :heavy_check_mark:                                                                 | The BYOK credential ID (UUID).                                                     | 11111111-2222-3333-4444-555555555555                                               |
| `updateBYOKKeyRequest`                                                             | [components.UpdateBYOKKeyRequest](../../models/components/updatebyokkeyrequest.md) | :heavy_check_mark:                                                                 | N/A                                                                                | {<br/>"disabled": false,<br/>"name": "Updated OpenAI Key"<br/>}                    |
| `opts`                                                                             | [][operations.Option](../../models/operations/option.md)                           | :heavy_minus_sign:                                                                 | The options for this request.                                                      |                                                                                    |

### Response

**[*components.UpdateBYOKKeyResponse](../../models/components/updatebyokkeyresponse.md), error**

### Errors

| Error Type                            | Status Code                           | Content Type                          |
| ------------------------------------- | ------------------------------------- | ------------------------------------- |
| sdkerrors.BadRequestResponseError     | 400                                   | application/json                      |
| sdkerrors.UnauthorizedResponseError   | 401                                   | application/json                      |
| sdkerrors.NotFoundResponseError       | 404                                   | application/json                      |
| sdkerrors.InternalServerResponseError | 500                                   | application/json                      |
| sdkerrors.APIError                    | 4XX, 5XX                              | \*/\*                                 |