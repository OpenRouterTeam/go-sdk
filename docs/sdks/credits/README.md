# Credits

## Overview

Credit management endpoints

### Available Operations

* [GetCredits](#getcredits) - Get remaining credits
* [CreateCoinbaseCharge](#createcoinbasecharge) - Create a Coinbase charge for crypto payment

## GetCredits

Get total credits purchased and used for the authenticated user. [Provisioning key](/docs/guides/overview/auth/provisioning-api-keys) required.

### Example Usage

<!-- UsageSnippet language="go" operationID="getCredits" method="get" path="/credits" -->
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

    res, err := s.Credits.GetCredits(ctx)
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

**[*operations.GetCreditsResponse](../../models/operations/getcreditsresponse.md), error**

### Errors

| Error Type                            | Status Code                           | Content Type                          |
| ------------------------------------- | ------------------------------------- | ------------------------------------- |
| sdkerrors.UnauthorizedResponseError   | 401                                   | application/json                      |
| sdkerrors.ForbiddenResponseError      | 403                                   | application/json                      |
| sdkerrors.InternalServerResponseError | 500                                   | application/json                      |
| sdkerrors.APIError                    | 4XX, 5XX                              | \*/\*                                 |

## CreateCoinbaseCharge

Create a Coinbase charge for crypto payment

### Example Usage

<!-- UsageSnippet language="go" operationID="createCoinbaseCharge" method="post" path="/credits/coinbase" -->
```go
package main

import(
	"context"
	openrouter "github.com/OpenRouterTeam/go-sdk"
	"os"
	"github.com/OpenRouterTeam/go-sdk/models/operations"
	"github.com/OpenRouterTeam/go-sdk/models/components"
	"log"
)

func main() {
    ctx := context.Background()

    s := openrouter.New()

    res, err := s.Credits.CreateCoinbaseCharge(ctx, operations.CreateCoinbaseChargeSecurity{
        Bearer: os.Getenv("OPENROUTER_BEARER"),
    }, components.CreateChargeRequest{
        Amount: 100,
        Sender: "0x1234567890123456789012345678901234567890",
        ChainID: components.ChainIDOne,
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

| Parameter                                                                                          | Type                                                                                               | Required                                                                                           | Description                                                                                        |
| -------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                              | [context.Context](https://pkg.go.dev/context#Context)                                              | :heavy_check_mark:                                                                                 | The context to use for the request.                                                                |
| `request`                                                                                          | [components.CreateChargeRequest](../../models/components/createchargerequest.md)                   | :heavy_check_mark:                                                                                 | The request object to use for the request.                                                         |
| `security`                                                                                         | [operations.CreateCoinbaseChargeSecurity](../../models/operations/createcoinbasechargesecurity.md) | :heavy_check_mark:                                                                                 | The security requirements to use for the request.                                                  |
| `opts`                                                                                             | [][operations.Option](../../models/operations/option.md)                                           | :heavy_minus_sign:                                                                                 | The options for this request.                                                                      |

### Response

**[*operations.CreateCoinbaseChargeResponse](../../models/operations/createcoinbasechargeresponse.md), error**

### Errors

| Error Type                             | Status Code                            | Content Type                           |
| -------------------------------------- | -------------------------------------- | -------------------------------------- |
| sdkerrors.BadRequestResponseError      | 400                                    | application/json                       |
| sdkerrors.UnauthorizedResponseError    | 401                                    | application/json                       |
| sdkerrors.TooManyRequestsResponseError | 429                                    | application/json                       |
| sdkerrors.InternalServerResponseError  | 500                                    | application/json                       |
| sdkerrors.APIError                     | 4XX, 5XX                               | \*/\*                                  |