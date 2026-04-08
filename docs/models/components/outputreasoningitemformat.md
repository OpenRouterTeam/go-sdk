# OutputReasoningItemFormat

The format of the reasoning content

## Example Usage

```go
import (
	"github.com/OpenRouterTeam/go-sdk/models/components"
)

value := components.OutputReasoningItemFormatUnknown

// Open enum: custom values can be created with a direct type cast
custom := components.OutputReasoningItemFormat("custom_value")
```


## Values

| Name                                              | Value                                             |
| ------------------------------------------------- | ------------------------------------------------- |
| `OutputReasoningItemFormatUnknown`                | unknown                                           |
| `OutputReasoningItemFormatOpenaiResponsesV1`      | openai-responses-v1                               |
| `OutputReasoningItemFormatAzureOpenaiResponsesV1` | azure-openai-responses-v1                         |
| `OutputReasoningItemFormatXaiResponsesV1`         | xai-responses-v1                                  |
| `OutputReasoningItemFormatAnthropicClaudeV1`      | anthropic-claude-v1                               |
| `OutputReasoningItemFormatGoogleGeminiV1`         | google-gemini-v1                                  |