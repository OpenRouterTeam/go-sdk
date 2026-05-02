# Stt

## Overview

Speech-to-text endpoints

### Available Operations

* [CreateTranscription](#createtranscription) - Create transcription

## CreateTranscription

Transcribes audio into text

### Example Usage

<!-- UsageSnippet language="go" operationID="createAudioTranscriptions" method="post" path="/audio/transcriptions" -->
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

    res, err := s.Stt.CreateTranscription(ctx, components.STTRequest{
        InputAudio: components.STTInputAudio{
            Data: "UklGRiQA...",
            Format: "wav",
        },
        Language: openrouter.Pointer("en"),
        Model: "openai/whisper-large-v3",
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

| Parameter                                                      | Type                                                           | Required                                                       | Description                                                    |
| -------------------------------------------------------------- | -------------------------------------------------------------- | -------------------------------------------------------------- | -------------------------------------------------------------- |
| `ctx`                                                          | [context.Context](https://pkg.go.dev/context#Context)          | :heavy_check_mark:                                             | The context to use for the request.                            |
| `request`                                                      | [components.STTRequest](../../models/components/sttrequest.md) | :heavy_check_mark:                                             | The request object to use for the request.                     |
| `opts`                                                         | [][operations.Option](../../models/operations/option.md)       | :heavy_minus_sign:                                             | The options for this request.                                  |

### Response

**[*components.STTResponse](../../models/components/sttresponse.md), error**

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