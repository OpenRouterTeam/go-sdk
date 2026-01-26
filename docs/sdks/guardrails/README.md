# Guardrails

## Overview

Guardrails endpoints

### Available Operations

* [List](#list) - List guardrails
* [Create](#create) - Create a guardrail
* [Get](#get) - Get a guardrail
* [Update](#update) - Update a guardrail
* [Delete](#delete) - Delete a guardrail
* [ListKeyAssignments](#listkeyassignments) - List all key assignments
* [ListMemberAssignments](#listmemberassignments) - List all member assignments
* [ListGuardrailKeyAssignments](#listguardrailkeyassignments) - List key assignments for a guardrail
* [BulkAssignKeys](#bulkassignkeys) - Bulk assign keys to a guardrail
* [ListGuardrailMemberAssignments](#listguardrailmemberassignments) - List member assignments for a guardrail
* [BulkAssignMembers](#bulkassignmembers) - Bulk assign members to a guardrail
* [BulkUnassignKeys](#bulkunassignkeys) - Bulk unassign keys from a guardrail
* [BulkUnassignMembers](#bulkunassignmembers) - Bulk unassign members from a guardrail

## List

List all guardrails for the authenticated user. [Provisioning key](/docs/guides/overview/auth/provisioning-api-keys) required.

### Example Usage

<!-- UsageSnippet language="go" operationID="listGuardrails" method="get" path="/guardrails" -->
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

    res, err := s.Guardrails.List(ctx, nil, nil)
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
| `offset`                                                 | **string*                                                | :heavy_minus_sign:                                       | Number of records to skip for pagination                 | 0                                                        |
| `limit`                                                  | **string*                                                | :heavy_minus_sign:                                       | Maximum number of records to return (max 100)            | 50                                                       |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |                                                          |

### Response

**[*operations.ListGuardrailsResponse](../../models/operations/listguardrailsresponse.md), error**

### Errors

| Error Type                            | Status Code                           | Content Type                          |
| ------------------------------------- | ------------------------------------- | ------------------------------------- |
| sdkerrors.UnauthorizedResponseError   | 401                                   | application/json                      |
| sdkerrors.InternalServerResponseError | 500                                   | application/json                      |
| sdkerrors.APIError                    | 4XX, 5XX                              | \*/\*                                 |

## Create

Create a new guardrail for the authenticated user. [Provisioning key](/docs/guides/overview/auth/provisioning-api-keys) required.

### Example Usage

<!-- UsageSnippet language="go" operationID="createGuardrail" method="post" path="/guardrails" -->
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

    res, err := s.Guardrails.Create(ctx, operations.CreateGuardrailRequest{
        Name: "My New Guardrail",
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
| `request`                                                                              | [operations.CreateGuardrailRequest](../../models/operations/createguardrailrequest.md) | :heavy_check_mark:                                                                     | The request object to use for the request.                                             |
| `opts`                                                                                 | [][operations.Option](../../models/operations/option.md)                               | :heavy_minus_sign:                                                                     | The options for this request.                                                          |

### Response

**[*operations.CreateGuardrailResponse](../../models/operations/createguardrailresponse.md), error**

### Errors

| Error Type                            | Status Code                           | Content Type                          |
| ------------------------------------- | ------------------------------------- | ------------------------------------- |
| sdkerrors.BadRequestResponseError     | 400                                   | application/json                      |
| sdkerrors.UnauthorizedResponseError   | 401                                   | application/json                      |
| sdkerrors.InternalServerResponseError | 500                                   | application/json                      |
| sdkerrors.APIError                    | 4XX, 5XX                              | \*/\*                                 |

## Get

Get a single guardrail by ID. [Provisioning key](/docs/guides/overview/auth/provisioning-api-keys) required.

### Example Usage

<!-- UsageSnippet language="go" operationID="getGuardrail" method="get" path="/guardrails/{id}" -->
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

    res, err := s.Guardrails.Get(ctx, "550e8400-e29b-41d4-a716-446655440000")
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
| `id`                                                     | *string*                                                 | :heavy_check_mark:                                       | The unique identifier of the guardrail to retrieve       | 550e8400-e29b-41d4-a716-446655440000                     |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |                                                          |

### Response

**[*operations.GetGuardrailResponse](../../models/operations/getguardrailresponse.md), error**

### Errors

| Error Type                            | Status Code                           | Content Type                          |
| ------------------------------------- | ------------------------------------- | ------------------------------------- |
| sdkerrors.UnauthorizedResponseError   | 401                                   | application/json                      |
| sdkerrors.NotFoundResponseError       | 404                                   | application/json                      |
| sdkerrors.InternalServerResponseError | 500                                   | application/json                      |
| sdkerrors.APIError                    | 4XX, 5XX                              | \*/\*                                 |

## Update

Update an existing guardrail. [Provisioning key](/docs/guides/overview/auth/provisioning-api-keys) required.

### Example Usage

<!-- UsageSnippet language="go" operationID="updateGuardrail" method="patch" path="/guardrails/{id}" -->
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

    res, err := s.Guardrails.Update(ctx, "550e8400-e29b-41d4-a716-446655440000", operations.UpdateGuardrailRequestBody{})
    if err != nil {
        log.Fatal(err)
    }
    if res != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                               | Type                                                                                                                    | Required                                                                                                                | Description                                                                                                             | Example                                                                                                                 |
| ----------------------------------------------------------------------------------------------------------------------- | ----------------------------------------------------------------------------------------------------------------------- | ----------------------------------------------------------------------------------------------------------------------- | ----------------------------------------------------------------------------------------------------------------------- | ----------------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                                   | [context.Context](https://pkg.go.dev/context#Context)                                                                   | :heavy_check_mark:                                                                                                      | The context to use for the request.                                                                                     |                                                                                                                         |
| `id`                                                                                                                    | *string*                                                                                                                | :heavy_check_mark:                                                                                                      | The unique identifier of the guardrail to update                                                                        | 550e8400-e29b-41d4-a716-446655440000                                                                                    |
| `requestBody`                                                                                                           | [operations.UpdateGuardrailRequestBody](../../models/operations/updateguardrailrequestbody.md)                          | :heavy_check_mark:                                                                                                      | N/A                                                                                                                     | {<br/>"name": "Updated Guardrail Name",<br/>"description": "Updated description",<br/>"limit_usd": 75,<br/>"reset_interval": "weekly"<br/>} |
| `opts`                                                                                                                  | [][operations.Option](../../models/operations/option.md)                                                                | :heavy_minus_sign:                                                                                                      | The options for this request.                                                                                           |                                                                                                                         |

### Response

**[*operations.UpdateGuardrailResponse](../../models/operations/updateguardrailresponse.md), error**

### Errors

| Error Type                            | Status Code                           | Content Type                          |
| ------------------------------------- | ------------------------------------- | ------------------------------------- |
| sdkerrors.BadRequestResponseError     | 400                                   | application/json                      |
| sdkerrors.UnauthorizedResponseError   | 401                                   | application/json                      |
| sdkerrors.NotFoundResponseError       | 404                                   | application/json                      |
| sdkerrors.InternalServerResponseError | 500                                   | application/json                      |
| sdkerrors.APIError                    | 4XX, 5XX                              | \*/\*                                 |

## Delete

Delete an existing guardrail. [Provisioning key](/docs/guides/overview/auth/provisioning-api-keys) required.

### Example Usage

<!-- UsageSnippet language="go" operationID="deleteGuardrail" method="delete" path="/guardrails/{id}" -->
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

    res, err := s.Guardrails.Delete(ctx, "550e8400-e29b-41d4-a716-446655440000")
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
| `id`                                                     | *string*                                                 | :heavy_check_mark:                                       | The unique identifier of the guardrail to delete         | 550e8400-e29b-41d4-a716-446655440000                     |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |                                                          |

### Response

**[*operations.DeleteGuardrailResponse](../../models/operations/deleteguardrailresponse.md), error**

### Errors

| Error Type                            | Status Code                           | Content Type                          |
| ------------------------------------- | ------------------------------------- | ------------------------------------- |
| sdkerrors.UnauthorizedResponseError   | 401                                   | application/json                      |
| sdkerrors.NotFoundResponseError       | 404                                   | application/json                      |
| sdkerrors.InternalServerResponseError | 500                                   | application/json                      |
| sdkerrors.APIError                    | 4XX, 5XX                              | \*/\*                                 |

## ListKeyAssignments

List all API key guardrail assignments for the authenticated user. [Provisioning key](/docs/guides/overview/auth/provisioning-api-keys) required.

### Example Usage

<!-- UsageSnippet language="go" operationID="listKeyAssignments" method="get" path="/guardrails/assignments/keys" -->
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

    res, err := s.Guardrails.ListKeyAssignments(ctx, nil, nil)
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
| `offset`                                                 | **string*                                                | :heavy_minus_sign:                                       | Number of records to skip for pagination                 | 0                                                        |
| `limit`                                                  | **string*                                                | :heavy_minus_sign:                                       | Maximum number of records to return (max 100)            | 50                                                       |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |                                                          |

### Response

**[*operations.ListKeyAssignmentsResponse](../../models/operations/listkeyassignmentsresponse.md), error**

### Errors

| Error Type                            | Status Code                           | Content Type                          |
| ------------------------------------- | ------------------------------------- | ------------------------------------- |
| sdkerrors.UnauthorizedResponseError   | 401                                   | application/json                      |
| sdkerrors.InternalServerResponseError | 500                                   | application/json                      |
| sdkerrors.APIError                    | 4XX, 5XX                              | \*/\*                                 |

## ListMemberAssignments

List all organization member guardrail assignments for the authenticated user. [Provisioning key](/docs/guides/overview/auth/provisioning-api-keys) required.

### Example Usage

<!-- UsageSnippet language="go" operationID="listMemberAssignments" method="get" path="/guardrails/assignments/members" -->
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

    res, err := s.Guardrails.ListMemberAssignments(ctx, nil, nil)
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
| `offset`                                                 | **string*                                                | :heavy_minus_sign:                                       | Number of records to skip for pagination                 | 0                                                        |
| `limit`                                                  | **string*                                                | :heavy_minus_sign:                                       | Maximum number of records to return (max 100)            | 50                                                       |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |                                                          |

### Response

**[*operations.ListMemberAssignmentsResponse](../../models/operations/listmemberassignmentsresponse.md), error**

### Errors

| Error Type                            | Status Code                           | Content Type                          |
| ------------------------------------- | ------------------------------------- | ------------------------------------- |
| sdkerrors.UnauthorizedResponseError   | 401                                   | application/json                      |
| sdkerrors.InternalServerResponseError | 500                                   | application/json                      |
| sdkerrors.APIError                    | 4XX, 5XX                              | \*/\*                                 |

## ListGuardrailKeyAssignments

List all API key assignments for a specific guardrail. [Provisioning key](/docs/guides/overview/auth/provisioning-api-keys) required.

### Example Usage

<!-- UsageSnippet language="go" operationID="listGuardrailKeyAssignments" method="get" path="/guardrails/{id}/assignments/keys" -->
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

    res, err := s.Guardrails.ListGuardrailKeyAssignments(ctx, "550e8400-e29b-41d4-a716-446655440000", nil, nil)
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
| `id`                                                     | *string*                                                 | :heavy_check_mark:                                       | The unique identifier of the guardrail                   | 550e8400-e29b-41d4-a716-446655440000                     |
| `offset`                                                 | **string*                                                | :heavy_minus_sign:                                       | Number of records to skip for pagination                 | 0                                                        |
| `limit`                                                  | **string*                                                | :heavy_minus_sign:                                       | Maximum number of records to return (max 100)            | 50                                                       |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |                                                          |

### Response

**[*operations.ListGuardrailKeyAssignmentsResponse](../../models/operations/listguardrailkeyassignmentsresponse.md), error**

### Errors

| Error Type                            | Status Code                           | Content Type                          |
| ------------------------------------- | ------------------------------------- | ------------------------------------- |
| sdkerrors.UnauthorizedResponseError   | 401                                   | application/json                      |
| sdkerrors.NotFoundResponseError       | 404                                   | application/json                      |
| sdkerrors.InternalServerResponseError | 500                                   | application/json                      |
| sdkerrors.APIError                    | 4XX, 5XX                              | \*/\*                                 |

## BulkAssignKeys

Assign multiple API keys to a specific guardrail. [Provisioning key](/docs/guides/overview/auth/provisioning-api-keys) required.

### Example Usage

<!-- UsageSnippet language="go" operationID="bulkAssignKeysToGuardrail" method="post" path="/guardrails/{id}/assignments/keys" -->
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

    res, err := s.Guardrails.BulkAssignKeys(ctx, "550e8400-e29b-41d4-a716-446655440000", operations.BulkAssignKeysToGuardrailRequestBody{
        KeyHashes: []string{
            "c56454edb818d6b14bc0d61c46025f1450b0f4012d12304ab40aacb519fcbc93",
        },
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

| Parameter                                                                                                          | Type                                                                                                               | Required                                                                                                           | Description                                                                                                        | Example                                                                                                            |
| ------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                                              | [context.Context](https://pkg.go.dev/context#Context)                                                              | :heavy_check_mark:                                                                                                 | The context to use for the request.                                                                                |                                                                                                                    |
| `id`                                                                                                               | *string*                                                                                                           | :heavy_check_mark:                                                                                                 | The unique identifier of the guardrail                                                                             | 550e8400-e29b-41d4-a716-446655440000                                                                               |
| `requestBody`                                                                                                      | [operations.BulkAssignKeysToGuardrailRequestBody](../../models/operations/bulkassignkeystoguardrailrequestbody.md) | :heavy_check_mark:                                                                                                 | N/A                                                                                                                |                                                                                                                    |
| `opts`                                                                                                             | [][operations.Option](../../models/operations/option.md)                                                           | :heavy_minus_sign:                                                                                                 | The options for this request.                                                                                      |                                                                                                                    |

### Response

**[*operations.BulkAssignKeysToGuardrailResponse](../../models/operations/bulkassignkeystoguardrailresponse.md), error**

### Errors

| Error Type                            | Status Code                           | Content Type                          |
| ------------------------------------- | ------------------------------------- | ------------------------------------- |
| sdkerrors.BadRequestResponseError     | 400                                   | application/json                      |
| sdkerrors.UnauthorizedResponseError   | 401                                   | application/json                      |
| sdkerrors.NotFoundResponseError       | 404                                   | application/json                      |
| sdkerrors.InternalServerResponseError | 500                                   | application/json                      |
| sdkerrors.APIError                    | 4XX, 5XX                              | \*/\*                                 |

## ListGuardrailMemberAssignments

List all organization member assignments for a specific guardrail. [Provisioning key](/docs/guides/overview/auth/provisioning-api-keys) required.

### Example Usage

<!-- UsageSnippet language="go" operationID="listGuardrailMemberAssignments" method="get" path="/guardrails/{id}/assignments/members" -->
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

    res, err := s.Guardrails.ListGuardrailMemberAssignments(ctx, "550e8400-e29b-41d4-a716-446655440000", nil, nil)
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
| `id`                                                     | *string*                                                 | :heavy_check_mark:                                       | The unique identifier of the guardrail                   | 550e8400-e29b-41d4-a716-446655440000                     |
| `offset`                                                 | **string*                                                | :heavy_minus_sign:                                       | Number of records to skip for pagination                 | 0                                                        |
| `limit`                                                  | **string*                                                | :heavy_minus_sign:                                       | Maximum number of records to return (max 100)            | 50                                                       |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |                                                          |

### Response

**[*operations.ListGuardrailMemberAssignmentsResponse](../../models/operations/listguardrailmemberassignmentsresponse.md), error**

### Errors

| Error Type                            | Status Code                           | Content Type                          |
| ------------------------------------- | ------------------------------------- | ------------------------------------- |
| sdkerrors.UnauthorizedResponseError   | 401                                   | application/json                      |
| sdkerrors.NotFoundResponseError       | 404                                   | application/json                      |
| sdkerrors.InternalServerResponseError | 500                                   | application/json                      |
| sdkerrors.APIError                    | 4XX, 5XX                              | \*/\*                                 |

## BulkAssignMembers

Assign multiple organization members to a specific guardrail. [Provisioning key](/docs/guides/overview/auth/provisioning-api-keys) required.

### Example Usage

<!-- UsageSnippet language="go" operationID="bulkAssignMembersToGuardrail" method="post" path="/guardrails/{id}/assignments/members" -->
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

    res, err := s.Guardrails.BulkAssignMembers(ctx, "550e8400-e29b-41d4-a716-446655440000", operations.BulkAssignMembersToGuardrailRequestBody{
        MemberUserIds: []string{
            "user_abc123",
            "user_def456",
        },
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

| Parameter                                                                                                                | Type                                                                                                                     | Required                                                                                                                 | Description                                                                                                              | Example                                                                                                                  |
| ------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                                                    | [context.Context](https://pkg.go.dev/context#Context)                                                                    | :heavy_check_mark:                                                                                                       | The context to use for the request.                                                                                      |                                                                                                                          |
| `id`                                                                                                                     | *string*                                                                                                                 | :heavy_check_mark:                                                                                                       | The unique identifier of the guardrail                                                                                   | 550e8400-e29b-41d4-a716-446655440000                                                                                     |
| `requestBody`                                                                                                            | [operations.BulkAssignMembersToGuardrailRequestBody](../../models/operations/bulkassignmemberstoguardrailrequestbody.md) | :heavy_check_mark:                                                                                                       | N/A                                                                                                                      |                                                                                                                          |
| `opts`                                                                                                                   | [][operations.Option](../../models/operations/option.md)                                                                 | :heavy_minus_sign:                                                                                                       | The options for this request.                                                                                            |                                                                                                                          |

### Response

**[*operations.BulkAssignMembersToGuardrailResponse](../../models/operations/bulkassignmemberstoguardrailresponse.md), error**

### Errors

| Error Type                            | Status Code                           | Content Type                          |
| ------------------------------------- | ------------------------------------- | ------------------------------------- |
| sdkerrors.BadRequestResponseError     | 400                                   | application/json                      |
| sdkerrors.UnauthorizedResponseError   | 401                                   | application/json                      |
| sdkerrors.NotFoundResponseError       | 404                                   | application/json                      |
| sdkerrors.InternalServerResponseError | 500                                   | application/json                      |
| sdkerrors.APIError                    | 4XX, 5XX                              | \*/\*                                 |

## BulkUnassignKeys

Unassign multiple API keys from a specific guardrail. [Provisioning key](/docs/guides/overview/auth/provisioning-api-keys) required.

### Example Usage

<!-- UsageSnippet language="go" operationID="bulkUnassignKeysFromGuardrail" method="post" path="/guardrails/{id}/assignments/keys/remove" -->
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

    res, err := s.Guardrails.BulkUnassignKeys(ctx, "550e8400-e29b-41d4-a716-446655440000", operations.BulkUnassignKeysFromGuardrailRequestBody{
        KeyHashes: []string{
            "c56454edb818d6b14bc0d61c46025f1450b0f4012d12304ab40aacb519fcbc93",
        },
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

| Parameter                                                                                                                  | Type                                                                                                                       | Required                                                                                                                   | Description                                                                                                                | Example                                                                                                                    |
| -------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                                      | [context.Context](https://pkg.go.dev/context#Context)                                                                      | :heavy_check_mark:                                                                                                         | The context to use for the request.                                                                                        |                                                                                                                            |
| `id`                                                                                                                       | *string*                                                                                                                   | :heavy_check_mark:                                                                                                         | The unique identifier of the guardrail                                                                                     | 550e8400-e29b-41d4-a716-446655440000                                                                                       |
| `requestBody`                                                                                                              | [operations.BulkUnassignKeysFromGuardrailRequestBody](../../models/operations/bulkunassignkeysfromguardrailrequestbody.md) | :heavy_check_mark:                                                                                                         | N/A                                                                                                                        |                                                                                                                            |
| `opts`                                                                                                                     | [][operations.Option](../../models/operations/option.md)                                                                   | :heavy_minus_sign:                                                                                                         | The options for this request.                                                                                              |                                                                                                                            |

### Response

**[*operations.BulkUnassignKeysFromGuardrailResponse](../../models/operations/bulkunassignkeysfromguardrailresponse.md), error**

### Errors

| Error Type                            | Status Code                           | Content Type                          |
| ------------------------------------- | ------------------------------------- | ------------------------------------- |
| sdkerrors.BadRequestResponseError     | 400                                   | application/json                      |
| sdkerrors.UnauthorizedResponseError   | 401                                   | application/json                      |
| sdkerrors.NotFoundResponseError       | 404                                   | application/json                      |
| sdkerrors.InternalServerResponseError | 500                                   | application/json                      |
| sdkerrors.APIError                    | 4XX, 5XX                              | \*/\*                                 |

## BulkUnassignMembers

Unassign multiple organization members from a specific guardrail. [Provisioning key](/docs/guides/overview/auth/provisioning-api-keys) required.

### Example Usage

<!-- UsageSnippet language="go" operationID="bulkUnassignMembersFromGuardrail" method="post" path="/guardrails/{id}/assignments/members/remove" -->
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

    res, err := s.Guardrails.BulkUnassignMembers(ctx, "550e8400-e29b-41d4-a716-446655440000", operations.BulkUnassignMembersFromGuardrailRequestBody{
        MemberUserIds: []string{
            "user_abc123",
            "user_def456",
        },
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

| Parameter                                                                                                                        | Type                                                                                                                             | Required                                                                                                                         | Description                                                                                                                      | Example                                                                                                                          |
| -------------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                                            | [context.Context](https://pkg.go.dev/context#Context)                                                                            | :heavy_check_mark:                                                                                                               | The context to use for the request.                                                                                              |                                                                                                                                  |
| `id`                                                                                                                             | *string*                                                                                                                         | :heavy_check_mark:                                                                                                               | The unique identifier of the guardrail                                                                                           | 550e8400-e29b-41d4-a716-446655440000                                                                                             |
| `requestBody`                                                                                                                    | [operations.BulkUnassignMembersFromGuardrailRequestBody](../../models/operations/bulkunassignmembersfromguardrailrequestbody.md) | :heavy_check_mark:                                                                                                               | N/A                                                                                                                              |                                                                                                                                  |
| `opts`                                                                                                                           | [][operations.Option](../../models/operations/option.md)                                                                         | :heavy_minus_sign:                                                                                                               | The options for this request.                                                                                                    |                                                                                                                                  |

### Response

**[*operations.BulkUnassignMembersFromGuardrailResponse](../../models/operations/bulkunassignmembersfromguardrailresponse.md), error**

### Errors

| Error Type                            | Status Code                           | Content Type                          |
| ------------------------------------- | ------------------------------------- | ------------------------------------- |
| sdkerrors.BadRequestResponseError     | 400                                   | application/json                      |
| sdkerrors.UnauthorizedResponseError   | 401                                   | application/json                      |
| sdkerrors.NotFoundResponseError       | 404                                   | application/json                      |
| sdkerrors.InternalServerResponseError | 500                                   | application/json                      |
| sdkerrors.APIError                    | 4XX, 5XX                              | \*/\*                                 |