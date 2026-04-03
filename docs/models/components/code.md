# Code

## Example Usage

```go
import (
	"github.com/OpenRouterTeam/go-sdk/models/components"
)

value := components.CodeServerError

// Open enum: custom values can be created with a direct type cast
custom := components.Code("custom_value")
```


## Values

| Name                              | Value                             |
| --------------------------------- | --------------------------------- |
| `CodeServerError`                 | server_error                      |
| `CodeRateLimitExceeded`           | rate_limit_exceeded               |
| `CodeInvalidPrompt`               | invalid_prompt                    |
| `CodeVectorStoreTimeout`          | vector_store_timeout              |
| `CodeInvalidImage`                | invalid_image                     |
| `CodeInvalidImageFormat`          | invalid_image_format              |
| `CodeInvalidBase64Image`          | invalid_base64_image              |
| `CodeInvalidImageURL`             | invalid_image_url                 |
| `CodeImageTooLarge`               | image_too_large                   |
| `CodeImageTooSmall`               | image_too_small                   |
| `CodeImageParseError`             | image_parse_error                 |
| `CodeImageContentPolicyViolation` | image_content_policy_violation    |
| `CodeInvalidImageMode`            | invalid_image_mode                |
| `CodeImageFileTooLarge`           | image_file_too_large              |
| `CodeUnsupportedImageMediaType`   | unsupported_image_media_type      |
| `CodeEmptyImageFile`              | empty_image_file                  |
| `CodeFailedToDownloadImage`       | failed_to_download_image          |
| `CodeImageFileNotFound`           | image_file_not_found              |