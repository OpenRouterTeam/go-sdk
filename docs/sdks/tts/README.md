# Tts

## Overview

Text-to-speech endpoints

### Available Operations

* [CreateSpeech](#createspeech) - Create speech

## CreateSpeech

Synthesizes audio from the input text

### Example Usage

<!-- UsageSnippet language="go" operationID="createAudioSpeech" method="post" path="/audio/speech" -->
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

    res, err := s.Tts.CreateSpeech(ctx, components.SpeechRequest{
        Input: "Hello world",
        Model: "elevenlabs/eleven-turbo-v2",
        Speed: openrouter.Pointer[float64](1.0),
        Voice: "alloy",
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

| Parameter                                                            | Type                                                                 | Required                                                             | Description                                                          |
| -------------------------------------------------------------------- | -------------------------------------------------------------------- | -------------------------------------------------------------------- | -------------------------------------------------------------------- |
| `ctx`                                                                | [context.Context](https://pkg.go.dev/context#Context)                | :heavy_check_mark:                                                   | The context to use for the request.                                  |
| `request`                                                            | [components.SpeechRequest](../../models/components/speechrequest.md) | :heavy_check_mark:                                                   | The request object to use for the request.                           |
| `opts`                                                               | [][operations.Option](../../models/operations/option.md)             | :heavy_minus_sign:                                                   | The options for this request.                                        |

### Response

**[io.ReadCloser](../../.md), error**

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