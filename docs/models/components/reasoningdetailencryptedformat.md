# ReasoningDetailEncryptedFormat

## Example Usage

```go
import (
	"github.com/OpenRouterTeam/go-sdk/models/components"
)

value := components.ReasoningDetailEncryptedFormatUnknown

// Open enum: custom values can be created with a direct type cast
custom := components.ReasoningDetailEncryptedFormat("custom_value")
```


## Values

| Name                                                   | Value                                                  |
| ------------------------------------------------------ | ------------------------------------------------------ |
| `ReasoningDetailEncryptedFormatUnknown`                | unknown                                                |
| `ReasoningDetailEncryptedFormatOpenaiResponsesV1`      | openai-responses-v1                                    |
| `ReasoningDetailEncryptedFormatAzureOpenaiResponsesV1` | azure-openai-responses-v1                              |
| `ReasoningDetailEncryptedFormatXaiResponsesV1`         | xai-responses-v1                                       |
| `ReasoningDetailEncryptedFormatAnthropicClaudeV1`      | anthropic-claude-v1                                    |
| `ReasoningDetailEncryptedFormatGoogleGeminiV1`         | google-gemini-v1                                       |