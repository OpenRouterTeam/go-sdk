# ReasoningItemFormat

## Example Usage

```go
import (
	"github.com/OpenRouterTeam/go-sdk/models/components"
)

value := components.ReasoningItemFormatUnknown

// Open enum: custom values can be created with a direct type cast
custom := components.ReasoningItemFormat("custom_value")
```


## Values

| Name                                        | Value                                       |
| ------------------------------------------- | ------------------------------------------- |
| `ReasoningItemFormatUnknown`                | unknown                                     |
| `ReasoningItemFormatOpenaiResponsesV1`      | openai-responses-v1                         |
| `ReasoningItemFormatAzureOpenaiResponsesV1` | azure-openai-responses-v1                   |
| `ReasoningItemFormatXaiResponsesV1`         | xai-responses-v1                            |
| `ReasoningItemFormatAnthropicClaudeV1`      | anthropic-claude-v1                         |
| `ReasoningItemFormatGoogleGeminiV1`         | google-gemini-v1                            |