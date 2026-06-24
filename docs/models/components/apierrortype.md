# APIErrorType

Canonical OpenRouter error type, stable across all API formats

## Example Usage

```go
import (
	"github.com/OpenRouterTeam/go-sdk/models/components"
)

value := components.APIErrorTypeContextLengthExceeded

// Open enum: custom values can be created with a direct type cast
custom := components.APIErrorType("custom_value")
```


## Values

| Name                                 | Value                                |
| ------------------------------------ | ------------------------------------ |
| `APIErrorTypeContextLengthExceeded`  | context_length_exceeded              |
| `APIErrorTypeMaxTokensExceeded`      | max_tokens_exceeded                  |
| `APIErrorTypeTokenLimitExceeded`     | token_limit_exceeded                 |
| `APIErrorTypeStringTooLong`          | string_too_long                      |
| `APIErrorTypeAuthentication`         | authentication                       |
| `APIErrorTypePermissionDenied`       | permission_denied                    |
| `APIErrorTypePaymentRequired`        | payment_required                     |
| `APIErrorTypeRateLimitExceeded`      | rate_limit_exceeded                  |
| `APIErrorTypeProviderOverloaded`     | provider_overloaded                  |
| `APIErrorTypeProviderUnavailable`    | provider_unavailable                 |
| `APIErrorTypeInvalidRequest`         | invalid_request                      |
| `APIErrorTypeInvalidPrompt`          | invalid_prompt                       |
| `APIErrorTypeNotFound`               | not_found                            |
| `APIErrorTypePreconditionFailed`     | precondition_failed                  |
| `APIErrorTypePayloadTooLarge`        | payload_too_large                    |
| `APIErrorTypeUnprocessable`          | unprocessable                        |
| `APIErrorTypeContentPolicyViolation` | content_policy_violation             |
| `APIErrorTypeRefusal`                | refusal                              |
| `APIErrorTypeInvalidImage`           | invalid_image                        |
| `APIErrorTypeImageTooLarge`          | image_too_large                      |
| `APIErrorTypeImageTooSmall`          | image_too_small                      |
| `APIErrorTypeUnsupportedImageFormat` | unsupported_image_format             |
| `APIErrorTypeImageNotFound`          | image_not_found                      |
| `APIErrorTypeImageDownloadFailed`    | image_download_failed                |
| `APIErrorTypeServer`                 | server                               |
| `APIErrorTypeTimeout`                | timeout                              |
| `APIErrorTypeUnmapped`               | unmapped                             |