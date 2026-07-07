# OpenRouter Go SDK

The [OpenRouter SDK](https://openrouter.ai/docs/sdks/go/api-reference/chat) is a Go client for building AI-powered features with OpenRouter. It gives you type-safe access to 400+ models across providers through an OpenAI-compatible API, plus OpenRouter-specific features like provider routing, guardrails, and analytics.

To learn more, see the [API Reference](https://openrouter.ai/docs/sdks/go/api-reference) and [Documentation](https://openrouter.ai/docs/sdks/go/api-reference/chat).

[![Built by Speakeasy](https://img.shields.io/badge/Built_by-SPEAKEASY-374151?style=for-the-badge&labelColor=f3f4f6)](https://www.speakeasy.com/?utm_source=openrouter&utm_campaign=go)
[![License: Apache 2.0](https://img.shields.io/badge/License-Apache_2.0-blue.svg?style=for-the-badge&labelColor=eff6ff)](https://opensource.org/licenses/Apache-2.0)

> [!NOTE]
> This SDK is in **beta**. Pin to a specific version to avoid unexpected breaking changes:
>
> ```bash
> go get github.com/OpenRouterTeam/go-sdk@v0.5.13
> ```

<!-- No Summary [summary] -->

## Overview

The OpenRouter Go SDK wraps the [OpenRouter API](https://openrouter.ai/docs) with idiomatic Go types, retries, and error handling. For a longer introduction, see [OVERVIEW.md](OVERVIEW.md).

- **Chat completions** with streaming and non-streaming responses
- **Embeddings, rerank, TTS, and video generation**
- **Beta Responses API** for agent-style workflows
- **Platform APIs** for API keys, credits, models, providers, guardrails, workspaces, and analytics
- **Configurable retries**, custom HTTP clients, and typed API errors

Install the module with:

```bash
go get github.com/OpenRouterTeam/go-sdk
```

See [examples/README.md](examples/README.md) for runnable examples, starting with [examples/chat](examples/chat).

<!-- Start Table of Contents [toc] -->
## Table of Contents
<!-- $toc-max-depth=2 -->
* [OpenRouter Go SDK](#openrouter-go-sdk)
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

<!-- Start Requirements [requirements] -->
## Requirements

This SDK requires Go 1.25 or higher.
<!-- End Requirements [requirements] -->

<!-- Start SDK Example Usage [usage] -->
## SDK Example Usage

### Example

```go
package main

import (
	"context"
	openrouter "github.com/OpenRouterTeam/go-sdk"
	"log"
	"os"
)

func main() {
	ctx := context.Background()

	s := openrouter.New(
		openrouter.WithSecurity(os.Getenv("OPENROUTER_API_KEY")),
	)

	res, err := s.Analytics.GetUserActivity(ctx, nil, nil, nil)
	if err != nil {
		log.Fatal(err)
	}
	if res != nil {
		// handle response
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
	"log"
	"os"
)

func main() {
	ctx := context.Background()

	s := openrouter.New(
		openrouter.WithSecurity(os.Getenv("OPENROUTER_API_KEY")),
	)

	res, err := s.Analytics.GetUserActivity(ctx, nil, nil, nil)
	if err != nil {
		log.Fatal(err)
	}
	if res != nil {
		// handle response
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
	"github.com/OpenRouterTeam/go-sdk/models/operations"
	"log"
	"os"
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
<!-- End Authentication [security] -->

<!-- Start Available Resources and Operations [operations] -->
## Available Resources and Operations

<details open>
<summary>Available methods</summary>

### [Analytics](docs/sdks/analytics/README.md)

* [GetUserActivity](docs/sdks/analytics/README.md#getuseractivity) - Get user activity grouped by endpoint

### [APIKeys](docs/sdks/apikeys/README.md)

* [GetCurrentKeyMetadata](docs/sdks/apikeys/README.md#getcurrentkeymetadata) - Get current API key
* [List](docs/sdks/apikeys/README.md#list) - List API keys
* [Create](docs/sdks/apikeys/README.md#create) - Create a new API key
* [Delete](docs/sdks/apikeys/README.md#delete) - Delete an API key
* [Get](docs/sdks/apikeys/README.md#get) - Get a single API key
* [Update](docs/sdks/apikeys/README.md#update) - Update an API key

### [Benchmarks](docs/sdks/benchmarks/README.md)

* [GetBenchmarks](docs/sdks/benchmarks/README.md#getbenchmarks) - List Benchmarks

### [Beta.Analytics](docs/sdks/betaanalytics/README.md)

* [GetAnalyticsMeta](docs/sdks/betaanalytics/README.md#getanalyticsmeta) - Get available analytics metrics and dimensions
* [QueryAnalytics](docs/sdks/betaanalytics/README.md#queryanalytics) - Query analytics data

### [Beta.Responses](docs/sdks/responses/README.md)

* [Send](docs/sdks/responses/README.md#send) - Create a response

### [Byok](docs/sdks/byok/README.md)

* [List](docs/sdks/byok/README.md#list) - List BYOK provider credentials
* [Create](docs/sdks/byok/README.md#create) - Create a BYOK provider credential
* [Delete](docs/sdks/byok/README.md#delete) - Delete a BYOK provider credential
* [Get](docs/sdks/byok/README.md#get) - Get a BYOK provider credential
* [Update](docs/sdks/byok/README.md#update) - Update a BYOK provider credential

### [Chat](docs/sdks/chat/README.md)

* [Send](docs/sdks/chat/README.md#send) - Create a chat completion

### [Classifications](docs/sdks/classifications/README.md)

* [GetTaskClassifications](docs/sdks/classifications/README.md#gettaskclassifications) - Task classification market share

### [Credits](docs/sdks/credits/README.md)

* [GetCredits](docs/sdks/credits/README.md#getcredits) - Get remaining credits

### [Datasets](docs/sdks/datasets/README.md)

* [GetAppRankings](docs/sdks/datasets/README.md#getapprankings) - Top apps by token usage
* [GetRankingsDaily](docs/sdks/datasets/README.md#getrankingsdaily) - Daily token totals for top 50 models

### [Embeddings](docs/sdks/embeddings/README.md)

* [Generate](docs/sdks/embeddings/README.md#generate) - Submit an embedding request
* [ListModels](docs/sdks/embeddings/README.md#listmodels) - List all embeddings models

### [Endpoints](docs/sdks/endpoints/README.md)

* [ListZdrEndpoints](docs/sdks/endpoints/README.md#listzdrendpoints) - Preview the impact of ZDR on the available endpoints
* [List](docs/sdks/endpoints/README.md#list) - List all endpoints for a model

### [Files](docs/sdks/files/README.md)

* [List](docs/sdks/files/README.md#list) - List files
* [Upload](docs/sdks/files/README.md#upload) - Upload a file
* [Delete](docs/sdks/files/README.md#delete) - Delete a file
* [Retrieve](docs/sdks/files/README.md#retrieve) - Get file metadata
* [Download](docs/sdks/files/README.md#download) - Download file content

### [Generations](docs/sdks/generations/README.md)

* [GetGeneration](docs/sdks/generations/README.md#getgeneration) - Get request & usage metadata for a generation
* [ListGenerationContent](docs/sdks/generations/README.md#listgenerationcontent) - Get stored prompt and completion content for a generation

### [Guardrails](docs/sdks/guardrails/README.md)

* [List](docs/sdks/guardrails/README.md#list) - List guardrails
* [Create](docs/sdks/guardrails/README.md#create) - Create a guardrail
* [Delete](docs/sdks/guardrails/README.md#delete) - Delete a guardrail
* [Get](docs/sdks/guardrails/README.md#get) - Get a guardrail
* [Update](docs/sdks/guardrails/README.md#update) - Update a guardrail
* [ListGuardrailKeyAssignments](docs/sdks/guardrails/README.md#listguardrailkeyassignments) - List key assignments for a guardrail
* [BulkAssignKeys](docs/sdks/guardrails/README.md#bulkassignkeys) - Bulk assign keys to a guardrail
* [BulkUnassignKeys](docs/sdks/guardrails/README.md#bulkunassignkeys) - Bulk unassign keys from a guardrail
* [ListGuardrailMemberAssignments](docs/sdks/guardrails/README.md#listguardrailmemberassignments) - List member assignments for a guardrail
* [BulkAssignMembers](docs/sdks/guardrails/README.md#bulkassignmembers) - Bulk assign members to a guardrail
* [BulkUnassignMembers](docs/sdks/guardrails/README.md#bulkunassignmembers) - Bulk unassign members from a guardrail
* [ListKeyAssignments](docs/sdks/guardrails/README.md#listkeyassignments) - List all key assignments
* [ListMemberAssignments](docs/sdks/guardrails/README.md#listmemberassignments) - List all member assignments

### [Images](docs/sdks/images/README.md)

* [Generate](docs/sdks/images/README.md#generate) - Generate an image
* [ListModels](docs/sdks/images/README.md#listmodels) - List image generation models
* [ListModelEndpoints](docs/sdks/images/README.md#listmodelendpoints) - List endpoints for an image model

### [Models](docs/sdks/models/README.md)

* [Get](docs/sdks/models/README.md#get) - Get a model by its slug
* [List](docs/sdks/models/README.md#list) - List all models and their properties
* [Count](docs/sdks/models/README.md#count) - Get total count of available models
* [ListForUser](docs/sdks/models/README.md#listforuser) - List models filtered by user provider preferences, privacy settings, and guardrails

### [OAuth](docs/sdks/oauth/README.md)

* [ExchangeAuthCodeForAPIKey](docs/sdks/oauth/README.md#exchangeauthcodeforapikey) - Exchange authorization code for API key
* [CreateAuthCode](docs/sdks/oauth/README.md#createauthcode) - Create authorization code

### [Observability](docs/sdks/observability/README.md)

* [List](docs/sdks/observability/README.md#list) - List observability destinations
* [Create](docs/sdks/observability/README.md#create) - Create an observability destination
* [Delete](docs/sdks/observability/README.md#delete) - Delete an observability destination
* [Get](docs/sdks/observability/README.md#get) - Get an observability destination
* [Update](docs/sdks/observability/README.md#update) - Update an observability destination

### [Organization](docs/sdks/organization/README.md)

* [ListMembers](docs/sdks/organization/README.md#listmembers) - List organization members

### [Presets](docs/sdks/presets/README.md)

* [List](docs/sdks/presets/README.md#list) - List presets
* [Get](docs/sdks/presets/README.md#get) - Get a preset
* [CreatePresetsChatCompletions](docs/sdks/presets/README.md#createpresetschatcompletions) - Create a preset from a chat-completions request body
* [CreatePresetsMessages](docs/sdks/presets/README.md#createpresetsmessages) - Create a preset from a messages request body
* [CreatePresetsResponses](docs/sdks/presets/README.md#createpresetsresponses) - Create a preset from a responses request body
* [ListVersions](docs/sdks/presets/README.md#listversions) - List versions of a preset
* [GetVersion](docs/sdks/presets/README.md#getversion) - Get a specific version of a preset

### [Providers](docs/sdks/providers/README.md)

* [List](docs/sdks/providers/README.md#list) - List all providers

### [Rerank](docs/sdks/rerank/README.md)

* [Rerank](docs/sdks/rerank/README.md#rerank) - Submit a rerank request

### [Stt](docs/sdks/stt/README.md)

* [CreateTranscription](docs/sdks/stt/README.md#createtranscription) - Create transcription
* [CreateTranscriptionMultipart](docs/sdks/stt/README.md#createtranscriptionmultipart) - Create transcription

### [Tts](docs/sdks/tts/README.md)

* [CreateSpeech](docs/sdks/tts/README.md#createspeech) - Create speech

### [VideoGeneration](docs/sdks/videogeneration/README.md)

* [Generate](docs/sdks/videogeneration/README.md#generate) - Submit a video generation request
* [GetGeneration](docs/sdks/videogeneration/README.md#getgeneration) - Poll video generation status
* [GetVideoContent](docs/sdks/videogeneration/README.md#getvideocontent) - Download generated video content
* [ListVideosModels](docs/sdks/videogeneration/README.md#listvideosmodels) - List all video generation models

### [Workspaces](docs/sdks/workspaces/README.md)

* [List](docs/sdks/workspaces/README.md#list) - List workspaces
* [Create](docs/sdks/workspaces/README.md#create) - Create a workspace
* [Delete](docs/sdks/workspaces/README.md#delete) - Delete a workspace
* [Get](docs/sdks/workspaces/README.md#get) - Get a workspace
* [Update](docs/sdks/workspaces/README.md#update) - Update a workspace
* [ListBudgets](docs/sdks/workspaces/README.md#listbudgets) - List workspace budgets
* [DeleteBudget](docs/sdks/workspaces/README.md#deletebudget) - Delete a workspace budget
* [SetBudget](docs/sdks/workspaces/README.md#setbudget) - Create or update a workspace budget
* [ListMembers](docs/sdks/workspaces/README.md#listmembers) - List workspace members
* [BulkAddMembers](docs/sdks/workspaces/README.md#bulkaddmembers) - Bulk add members to a workspace
* [BulkRemoveMembers](docs/sdks/workspaces/README.md#bulkremovemembers) - Bulk remove members from a workspace

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

	res, err := s.Beta.Responses.Send(ctx, components.ResponsesRequest{
		Input: openrouter.Pointer(components.CreateInputsUnionStr(
			"Tell me a joke",
		)),
		Model: openrouter.Pointer("openai/gpt-4o"),
	}, components.MetadataLevelEnabled.ToPointer())
	if err != nil {
		log.Fatal(err)
	}
	if res != nil {
		defer res.ResponsesStreamingResponse.Close()

		for res.ResponsesStreamingResponse.Next() {
			event := res.ResponsesStreamingResponse.Value()
			log.Print(event)
			// Handle the event
		}
	}
}

```

[mdn-sse]: https://developer.mozilla.org/en-US/docs/Web/API/Server-sent_events/Using_server-sent_events
<!-- End Server-sent event streaming [eventstream] -->

<!-- Start Pagination [pagination] -->
## Pagination

Some of the endpoints in this SDK support pagination. To use pagination, you make your SDK calls as usual, but the
returned response object will have a `Next` method that can be called to pull down the next group of results. If the
return value of `Next` is `nil`, then there are no more pages to be fetched.

Here's an example of one such pagination call:
```go
package main

import (
	"context"
	openrouter "github.com/OpenRouterTeam/go-sdk"
	"github.com/OpenRouterTeam/go-sdk/optionalnullable"
	"log"
	"os"
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
<!-- End Pagination [pagination] -->

<!-- Start Retries [retries] -->
## Retries

Some of the endpoints in this SDK support retries. If you use the SDK without any configuration, it will fall back to the default retry strategy provided by the API. However, the default retry strategy can be overridden on a per-operation basis, or across the entire SDK.

To change the default retry strategy for a single API call, simply provide a `retry.Config` object to the call by using the `WithRetries` option:
```go
package main

import (
	"context"
	openrouter "github.com/OpenRouterTeam/go-sdk"
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

	res, err := s.Analytics.GetUserActivity(ctx, nil, nil, nil, operations.WithRetries(
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
		// handle response
	}
}

```

If you'd like to override the default retry strategy for all operations that support retries, you can use the `WithRetryConfig` option at SDK initialization:
```go
package main

import (
	"context"
	openrouter "github.com/OpenRouterTeam/go-sdk"
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

	res, err := s.Analytics.GetUserActivity(ctx, nil, nil, nil)
	if err != nil {
		log.Fatal(err)
	}
	if res != nil {
		// handle response
	}
}

```
<!-- End Retries [retries] -->

<!-- Start Error Handling [errors] -->
## Error Handling

Handling errors in this SDK should largely match your expectations. All operations return a response object or an error, they will never return both.

By Default, an API error will return `sdkerrors.APIError`. When custom error responses are specified for an operation, the SDK may also return their associated error. You can refer to respective *Errors* tables in SDK docs for more details on possible error types for each operation.

For example, the `GetUserActivity` function may return the following errors:

| Error Type                            | Status Code | Content Type     |
| ------------------------------------- | ----------- | ---------------- |
| sdkerrors.BadRequestResponseError     | 400         | application/json |
| sdkerrors.UnauthorizedResponseError   | 401         | application/json |
| sdkerrors.ForbiddenResponseError      | 403         | application/json |
| sdkerrors.NotFoundResponseError       | 404         | application/json |
| sdkerrors.InternalServerResponseError | 500         | application/json |
| sdkerrors.APIError                    | 4XX, 5XX    | \*/\*            |

### Example

```go
package main

import (
	"context"
	"errors"
	openrouter "github.com/OpenRouterTeam/go-sdk"
	"github.com/OpenRouterTeam/go-sdk/models/sdkerrors"
	"log"
	"os"
)

func main() {
	ctx := context.Background()

	s := openrouter.New(
		openrouter.WithSecurity(os.Getenv("OPENROUTER_API_KEY")),
	)

	res, err := s.Analytics.GetUserActivity(ctx, nil, nil, nil)
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

		var e *sdkerrors.ForbiddenResponseError
		if errors.As(err, &e) {
			// handle error
			log.Fatal(e.Error())
		}

		var e *sdkerrors.NotFoundResponseError
		if errors.As(err, &e) {
			// handle error
			log.Fatal(e.Error())
		}

		var e *sdkerrors.InternalServerResponseError
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
	"log"
	"os"
)

func main() {
	ctx := context.Background()

	s := openrouter.New(
		openrouter.WithServer("production"),
		openrouter.WithSecurity(os.Getenv("OPENROUTER_API_KEY")),
	)

	res, err := s.Analytics.GetUserActivity(ctx, nil, nil, nil)
	if err != nil {
		log.Fatal(err)
	}
	if res != nil {
		// handle response
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
	"log"
	"os"
)

func main() {
	ctx := context.Background()

	s := openrouter.New(
		openrouter.WithServerURL("https://openrouter.ai/api/v1"),
		openrouter.WithSecurity(os.Getenv("OPENROUTER_API_KEY")),
	)

	res, err := s.Analytics.GetUserActivity(ctx, nil, nil, nil)
	if err != nil {
		log.Fatal(err)
	}
	if res != nil {
		// handle response
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

This SDK is in beta. Breaking changes may ship in minor `0.x` releases. Pin to a specific module version in production, and review [RELEASES.md](RELEASES.md) before upgrading.

## Contributions

While we value open-source contributions to this SDK, this library is generated programmatically. Any manual changes added to internal files will be overwritten on the next generation.
We look forward to hearing your feedback. Feel free to open a PR or an issue with a proof of concept and we'll do our best to include it in a future release.

Safe to edit manually without being overwritten by Speakeasy generation:

- [USAGE.md](USAGE.md) — source for the README SDK Example Usage section
- [OVERVIEW.md](OVERVIEW.md) — extended introduction
- [examples/](examples/) — runnable example modules
- [docs/sdks/chat/README.md](docs/sdks/chat/README.md) and other files tracked in Speakeasy persistent edits

### SDK Created by [Speakeasy](https://www.speakeasy.com/?utm_source=openrouter&utm_campaign=go)
