# openrouter

Developer-friendly & type-safe Go SDK specifically catered to leverage *openrouter* API.

[![Built by Speakeasy](https://img.shields.io/badge/Built_by-SPEAKEASY-374151?style=for-the-badge&labelColor=f3f4f6)](https://www.speakeasy.com/?utm_source=openrouter&utm_campaign=go)
[![License: MIT](https://img.shields.io/badge/LICENSE_//_MIT-3b5bdb?style=for-the-badge&labelColor=eff6ff)](https://opensource.org/licenses/MIT)


<br /><br />
> [!IMPORTANT]
> This SDK is not yet ready for production use. To complete setup please follow the steps outlined in your [workspace](https://app.speakeasy.com/org/openrouter/sdk). Delete this section before > publishing to a package manager.

<!-- Start Summary [summary] -->
## Summary

OpenRouter API: OpenAI-compatible API with additional OpenRouter features

For more information about the API: [OpenRouter Documentation](https://openrouter.ai/docs)
<!-- End Summary [summary] -->

<!-- Start Table of Contents [toc] -->
## Table of Contents
<!-- $toc-max-depth=2 -->
* [openrouter](#openrouter)
  * [SDK Installation](#sdk-installation)
  * [SDK Example Usage](#sdk-example-usage)
  * [Authentication](#authentication)
  * [Available Resources and Operations](#available-resources-and-operations)
  * [Server-sent event streaming](#server-sent-event-streaming)
  * [Retries](#retries)
  * [Error Handling](#error-handling)
  * [Server Selection](#server-selection)
  * [Custom HTTP Client](#custom-http-client)
* [Development](#development)
  * [Maturity](#maturity)
  * [Contributions](#contributions)

<!-- End Table of Contents [toc] -->

<!-- Start SDK Installation [installation] -->
## SDK Installation

To add the SDK as a dependency to your project:
```bash
go get github.com/OpenRouterTeam/go-sdk
```
<!-- End SDK Installation [installation] -->

<!-- Start SDK Example Usage [usage] -->
## SDK Example Usage

### Example

```go
package main

import (
	"context"
	openrouter "github.com/OpenRouterTeam/go-sdk"
	"github.com/OpenRouterTeam/go-sdk/models/components"
	"log"
	"os"
)

func main() {
	ctx := context.Background()

	s := openrouter.New(
		openrouter.WithSecurity(os.Getenv("OPENROUTER_API_KEY")),
	)

	res, err := s.Beta.Responses.Send(ctx, components.OpenResponsesRequest{})
	if err != nil {
		log.Fatal(err)
	}
	if res != nil {
		defer res.Object.Close()

		for res.Object.Next() {
			event := res.Object.Value()
			log.Print(event)
			// Handle the event
		}
	}
}

```
<!-- End SDK Example Usage [usage] -->

<!-- Start Authentication [security] -->
## Authentication

### Per-Client Security Schemes

This SDK supports the following security scheme globally:

| Name     | Type | Scheme      | Environment Variable |
| -------- | ---- | ----------- | -------------------- |
| `APIKey` | http | HTTP Bearer | `OPENROUTER_API_KEY` |

You can configure it using the `WithSecurity` option when initializing the SDK client instance. For example:
```go
package main

import (
	"context"
	openrouter "github.com/OpenRouterTeam/go-sdk"
	"github.com/OpenRouterTeam/go-sdk/models/components"
	"log"
	"os"
)

func main() {
	ctx := context.Background()

	s := openrouter.New(
		openrouter.WithSecurity(os.Getenv("OPENROUTER_API_KEY")),
	)

	res, err := s.Beta.Responses.Send(ctx, components.OpenResponsesRequest{})
	if err != nil {
		log.Fatal(err)
	}
	if res != nil {
		defer res.Object.Close()

		for res.Object.Next() {
			event := res.Object.Value()
			log.Print(event)
			// Handle the event
		}
	}
}

```

### Per-Operation Security Schemes

Some operations in this SDK require the security scheme to be specified at the request level. For example:
```go
package main

import (
	"context"
	openrouter "github.com/OpenRouterTeam/go-sdk"
	"github.com/OpenRouterTeam/go-sdk/models/components"
	"github.com/OpenRouterTeam/go-sdk/models/operations"
	"log"
	"os"
)

func main() {
	ctx := context.Background()

	s := openrouter.New()

	res, err := s.Credits.CreateCoinbaseCharge(ctx, operations.CreateCoinbaseChargeSecurity{
		Bearer: os.Getenv("OPENROUTER_BEARER"),
	}, components.CreateChargeRequest{
		Amount:  100,
		Sender:  "0x1234567890123456789012345678901234567890",
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
<!-- End Authentication [security] -->

<!-- Start Available Resources and Operations [operations] -->
## Available Resources and Operations

<details open>
<summary>Available methods</summary>

### [Analytics](docs/sdks/analytics/README.md)

* [GetUserActivity](docs/sdks/analytics/README.md#getuseractivity) - Get user activity grouped by endpoint

### [APIKeys](docs/sdks/apikeys/README.md)

* [List](docs/sdks/apikeys/README.md#list) - List API keys
* [Create](docs/sdks/apikeys/README.md#create) - Create a new API key
* [Update](docs/sdks/apikeys/README.md#update) - Update an API key
* [Delete](docs/sdks/apikeys/README.md#delete) - Delete an API key
* [Get](docs/sdks/apikeys/README.md#get) - Get a single API key
* [GetCurrentKeyMetadata](docs/sdks/apikeys/README.md#getcurrentkeymetadata) - Get current API key

### [Beta.Responses](docs/sdks/responses/README.md)

* [Send](docs/sdks/responses/README.md#send) - Create a response

### [Chat](docs/sdks/chat/README.md)

* [Send](docs/sdks/chat/README.md#send) - Create a chat completion

### [Credits](docs/sdks/credits/README.md)

* [GetCredits](docs/sdks/credits/README.md#getcredits) - Get remaining credits
* [CreateCoinbaseCharge](docs/sdks/credits/README.md#createcoinbasecharge) - Create a Coinbase charge for crypto payment

### [Embeddings](docs/sdks/embeddings/README.md)

* [Generate](docs/sdks/embeddings/README.md#generate) - Submit an embedding request
* [ListModels](docs/sdks/embeddings/README.md#listmodels) - List all embeddings models

### [Endpoints](docs/sdks/endpoints/README.md)

* [List](docs/sdks/endpoints/README.md#list) - List all endpoints for a model
* [ListZdrEndpoints](docs/sdks/endpoints/README.md#listzdrendpoints) - Preview the impact of ZDR on the available endpoints

### [Generations](docs/sdks/generations/README.md)

* [GetGeneration](docs/sdks/generations/README.md#getgeneration) - Get request & usage metadata for a generation

### [Guardrails](docs/sdks/guardrails/README.md)

* [List](docs/sdks/guardrails/README.md#list) - List guardrails
* [Create](docs/sdks/guardrails/README.md#create) - Create a guardrail
* [Get](docs/sdks/guardrails/README.md#get) - Get a guardrail
* [Update](docs/sdks/guardrails/README.md#update) - Update a guardrail
* [Delete](docs/sdks/guardrails/README.md#delete) - Delete a guardrail
* [ListKeyAssignments](docs/sdks/guardrails/README.md#listkeyassignments) - List all key assignments
* [ListMemberAssignments](docs/sdks/guardrails/README.md#listmemberassignments) - List all member assignments
* [ListGuardrailKeyAssignments](docs/sdks/guardrails/README.md#listguardrailkeyassignments) - List key assignments for a guardrail
* [BulkAssignKeys](docs/sdks/guardrails/README.md#bulkassignkeys) - Bulk assign keys to a guardrail
* [ListGuardrailMemberAssignments](docs/sdks/guardrails/README.md#listguardrailmemberassignments) - List member assignments for a guardrail
* [BulkAssignMembers](docs/sdks/guardrails/README.md#bulkassignmembers) - Bulk assign members to a guardrail
* [BulkUnassignKeys](docs/sdks/guardrails/README.md#bulkunassignkeys) - Bulk unassign keys from a guardrail
* [BulkUnassignMembers](docs/sdks/guardrails/README.md#bulkunassignmembers) - Bulk unassign members from a guardrail

### [Models](docs/sdks/models/README.md)

* [Count](docs/sdks/models/README.md#count) - Get total count of available models
* [List](docs/sdks/models/README.md#list) - List all models and their properties
* [ListForUser](docs/sdks/models/README.md#listforuser) - List models filtered by user provider preferences, privacy settings, and guardrails

### [OAuth](docs/sdks/oauth/README.md)

* [ExchangeAuthCodeForAPIKey](docs/sdks/oauth/README.md#exchangeauthcodeforapikey) - Exchange authorization code for API key
* [CreateAuthCode](docs/sdks/oauth/README.md#createauthcode) - Create authorization code

### [Providers](docs/sdks/providers/README.md)

* [List](docs/sdks/providers/README.md#list) - List all providers

</details>
<!-- End Available Resources and Operations [operations] -->

<!-- Start Server-sent event streaming [eventstream] -->
## Server-sent event streaming

[Server-sent events][mdn-sse] are used to stream content from certain
operations. These operations will expose the stream as an iterable that
can be consumed using a simple `for` loop. The loop will
terminate when the server no longer has any events to send and closes the
underlying connection.

```go
package main

import (
	"context"
	openrouter "github.com/OpenRouterTeam/go-sdk"
	"github.com/OpenRouterTeam/go-sdk/models/components"
	"log"
	"os"
)

func main() {
	ctx := context.Background()

	s := openrouter.New(
		openrouter.WithSecurity(os.Getenv("OPENROUTER_API_KEY")),
	)

	res, err := s.Beta.Responses.Send(ctx, components.OpenResponsesRequest{})
	if err != nil {
		log.Fatal(err)
	}
	if res != nil {
		defer res.Object.Close()

		for res.Object.Next() {
			event := res.Object.Value()
			log.Print(event)
			// Handle the event
		}
	}
}

```

[mdn-sse]: https://developer.mozilla.org/en-US/docs/Web/API/Server-sent_events/Using_server-sent_events
<!-- End Server-sent event streaming [eventstream] -->

<!-- Start Retries [retries] -->
## Retries

Some of the endpoints in this SDK support retries. If you use the SDK without any configuration, it will fall back to the default retry strategy provided by the API. However, the default retry strategy can be overridden on a per-operation basis, or across the entire SDK.

To change the default retry strategy for a single API call, simply provide a `retry.Config` object to the call by using the `WithRetries` option:
```go
package main

import (
	"context"
	openrouter "github.com/OpenRouterTeam/go-sdk"
	"github.com/OpenRouterTeam/go-sdk/models/components"
	"github.com/OpenRouterTeam/go-sdk/retry"
	"log"
	"models/operations"
	"os"
)

func main() {
	ctx := context.Background()

	s := openrouter.New(
		openrouter.WithSecurity(os.Getenv("OPENROUTER_API_KEY")),
	)

	res, err := s.Beta.Responses.Send(ctx, components.OpenResponsesRequest{}, operations.WithRetries(
		retry.Config{
			Strategy: "backoff",
			Backoff: &retry.BackoffStrategy{
				InitialInterval: 1,
				MaxInterval:     50,
				Exponent:        1.1,
				MaxElapsedTime:  100,
			},
			RetryConnectionErrors: false,
		}))
	if err != nil {
		log.Fatal(err)
	}
	if res != nil {
		defer res.Object.Close()

		for res.Object.Next() {
			event := res.Object.Value()
			log.Print(event)
			// Handle the event
		}
	}
}

```

If you'd like to override the default retry strategy for all operations that support retries, you can use the `WithRetryConfig` option at SDK initialization:
```go
package main

import (
	"context"
	openrouter "github.com/OpenRouterTeam/go-sdk"
	"github.com/OpenRouterTeam/go-sdk/models/components"
	"github.com/OpenRouterTeam/go-sdk/retry"
	"log"
	"os"
)

func main() {
	ctx := context.Background()

	s := openrouter.New(
		openrouter.WithRetryConfig(
			retry.Config{
				Strategy: "backoff",
				Backoff: &retry.BackoffStrategy{
					InitialInterval: 1,
					MaxInterval:     50,
					Exponent:        1.1,
					MaxElapsedTime:  100,
				},
				RetryConnectionErrors: false,
			}),
		openrouter.WithSecurity(os.Getenv("OPENROUTER_API_KEY")),
	)

	res, err := s.Beta.Responses.Send(ctx, components.OpenResponsesRequest{})
	if err != nil {
		log.Fatal(err)
	}
	if res != nil {
		defer res.Object.Close()

		for res.Object.Next() {
			event := res.Object.Value()
			log.Print(event)
			// Handle the event
		}
	}
}

```
<!-- End Retries [retries] -->

<!-- Start Error Handling [errors] -->
## Error Handling

Handling errors in this SDK should largely match your expectations. All operations return a response object or an error, they will never return both.

By Default, an API error will return `sdkerrors.APIError`. When custom error responses are specified for an operation, the SDK may also return their associated error. You can refer to respective *Errors* tables in SDK docs for more details on possible error types for each operation.

For example, the `Send` function may return the following errors:

| Error Type                                 | Status Code | Content Type     |
| ------------------------------------------ | ----------- | ---------------- |
| sdkerrors.BadRequestResponseError          | 400         | application/json |
| sdkerrors.UnauthorizedResponseError        | 401         | application/json |
| sdkerrors.PaymentRequiredResponseError     | 402         | application/json |
| sdkerrors.NotFoundResponseError            | 404         | application/json |
| sdkerrors.RequestTimeoutResponseError      | 408         | application/json |
| sdkerrors.PayloadTooLargeResponseError     | 413         | application/json |
| sdkerrors.UnprocessableEntityResponseError | 422         | application/json |
| sdkerrors.TooManyRequestsResponseError     | 429         | application/json |
| sdkerrors.InternalServerResponseError      | 500         | application/json |
| sdkerrors.BadGatewayResponseError          | 502         | application/json |
| sdkerrors.ServiceUnavailableResponseError  | 503         | application/json |
| sdkerrors.EdgeNetworkTimeoutResponseError  | 524         | application/json |
| sdkerrors.ProviderOverloadedResponseError  | 529         | application/json |
| sdkerrors.APIError                         | 4XX, 5XX    | \*/\*            |

### Example

```go
package main

import (
	"context"
	"errors"
	openrouter "github.com/OpenRouterTeam/go-sdk"
	"github.com/OpenRouterTeam/go-sdk/models/components"
	"github.com/OpenRouterTeam/go-sdk/models/sdkerrors"
	"log"
	"os"
)

func main() {
	ctx := context.Background()

	s := openrouter.New(
		openrouter.WithSecurity(os.Getenv("OPENROUTER_API_KEY")),
	)

	res, err := s.Beta.Responses.Send(ctx, components.OpenResponsesRequest{})
	if err != nil {

		var e *sdkerrors.BadRequestResponseError
		if errors.As(err, &e) {
			// handle error
			log.Fatal(e.Error())
		}

		var e *sdkerrors.UnauthorizedResponseError
		if errors.As(err, &e) {
			// handle error
			log.Fatal(e.Error())
		}

		var e *sdkerrors.PaymentRequiredResponseError
		if errors.As(err, &e) {
			// handle error
			log.Fatal(e.Error())
		}

		var e *sdkerrors.NotFoundResponseError
		if errors.As(err, &e) {
			// handle error
			log.Fatal(e.Error())
		}

		var e *sdkerrors.RequestTimeoutResponseError
		if errors.As(err, &e) {
			// handle error
			log.Fatal(e.Error())
		}

		var e *sdkerrors.PayloadTooLargeResponseError
		if errors.As(err, &e) {
			// handle error
			log.Fatal(e.Error())
		}

		var e *sdkerrors.UnprocessableEntityResponseError
		if errors.As(err, &e) {
			// handle error
			log.Fatal(e.Error())
		}

		var e *sdkerrors.TooManyRequestsResponseError
		if errors.As(err, &e) {
			// handle error
			log.Fatal(e.Error())
		}

		var e *sdkerrors.InternalServerResponseError
		if errors.As(err, &e) {
			// handle error
			log.Fatal(e.Error())
		}

		var e *sdkerrors.BadGatewayResponseError
		if errors.As(err, &e) {
			// handle error
			log.Fatal(e.Error())
		}

		var e *sdkerrors.ServiceUnavailableResponseError
		if errors.As(err, &e) {
			// handle error
			log.Fatal(e.Error())
		}

		var e *sdkerrors.EdgeNetworkTimeoutResponseError
		if errors.As(err, &e) {
			// handle error
			log.Fatal(e.Error())
		}

		var e *sdkerrors.ProviderOverloadedResponseError
		if errors.As(err, &e) {
			// handle error
			log.Fatal(e.Error())
		}

		var e *sdkerrors.APIError
		if errors.As(err, &e) {
			// handle error
			log.Fatal(e.Error())
		}
	}
}

```
<!-- End Error Handling [errors] -->

<!-- Start Server Selection [server] -->
## Server Selection

### Select Server by Name

You can override the default server globally using the `WithServer(server string)` option when initializing the SDK client instance. The selected server will then be used as the default on the operations that use it. This table lists the names associated with the available servers:

| Name         | Server                         | Description       |
| ------------ | ------------------------------ | ----------------- |
| `production` | `https://openrouter.ai/api/v1` | Production server |

#### Example

```go
package main

import (
	"context"
	openrouter "github.com/OpenRouterTeam/go-sdk"
	"github.com/OpenRouterTeam/go-sdk/models/components"
	"log"
	"os"
)

func main() {
	ctx := context.Background()

	s := openrouter.New(
		openrouter.WithServer("production"),
		openrouter.WithSecurity(os.Getenv("OPENROUTER_API_KEY")),
	)

	res, err := s.Beta.Responses.Send(ctx, components.OpenResponsesRequest{})
	if err != nil {
		log.Fatal(err)
	}
	if res != nil {
		defer res.Object.Close()

		for res.Object.Next() {
			event := res.Object.Value()
			log.Print(event)
			// Handle the event
		}
	}
}

```

### Override Server URL Per-Client

The default server can also be overridden globally using the `WithServerURL(serverURL string)` option when initializing the SDK client instance. For example:
```go
package main

import (
	"context"
	openrouter "github.com/OpenRouterTeam/go-sdk"
	"github.com/OpenRouterTeam/go-sdk/models/components"
	"log"
	"os"
)

func main() {
	ctx := context.Background()

	s := openrouter.New(
		openrouter.WithServerURL("https://openrouter.ai/api/v1"),
		openrouter.WithSecurity(os.Getenv("OPENROUTER_API_KEY")),
	)

	res, err := s.Beta.Responses.Send(ctx, components.OpenResponsesRequest{})
	if err != nil {
		log.Fatal(err)
	}
	if res != nil {
		defer res.Object.Close()

		for res.Object.Next() {
			event := res.Object.Value()
			log.Print(event)
			// Handle the event
		}
	}
}

```
<!-- End Server Selection [server] -->

<!-- Start Custom HTTP Client [http-client] -->
## Custom HTTP Client

The Go SDK makes API calls that wrap an internal HTTP client. The requirements for the HTTP client are very simple. It must match this interface:

```go
type HTTPClient interface {
	Do(req *http.Request) (*http.Response, error)
}
```

The built-in `net/http` client satisfies this interface and a default client based on the built-in is provided by default. To replace this default with a client of your own, you can implement this interface yourself or provide your own client configured as desired. Here's a simple example, which adds a client with a 30 second timeout.

```go
import (
	"net/http"
	"time"

	"github.com/OpenRouterTeam/go-sdk"
)

var (
	httpClient = &http.Client{Timeout: 30 * time.Second}
	sdkClient  = openrouter.New(openrouter.WithClient(httpClient))
)
```

This can be a convenient way to configure timeouts, cookies, proxies, custom headers, and other low-level configuration.
<!-- End Custom HTTP Client [http-client] -->

<!-- Placeholder for Future Speakeasy SDK Sections -->

# Development

## Maturity

This SDK is in beta, and there may be breaking changes between versions without a major version update. Therefore, we recommend pinning usage
to a specific package version. This way, you can install the same version each time without breaking changes unless you are intentionally
looking for the latest version.

## Contributions

While we value open-source contributions to this SDK, this library is generated programmatically. Any manual changes added to internal files will be overwritten on the next generation. 
We look forward to hearing your feedback. Feel free to open a PR or an issue with a proof of concept and we'll do our best to include it in a future release. 

### SDK Created by [Speakeasy](https://www.speakeasy.com/?utm_source=openrouter&utm_campaign=go)
