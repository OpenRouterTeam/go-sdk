# ReasoningFormat

## Example Usage

```go
import (
	"github.com/OpenRouterTeam/go-sdk/models/components"
)

value := components.ReasoningFormatUnknown

// Open enum: custom values can be created with a direct type cast
custom := components.ReasoningFormat("custom_value")
```


## Values

| Name                                    | Value                                   |
| --------------------------------------- | --------------------------------------- |
| `ReasoningFormatUnknown`                | unknown                                 |
| `ReasoningFormatOpenaiResponsesV1`      | openai-responses-v1                     |
| `ReasoningFormatAzureOpenaiResponsesV1` | azure-openai-responses-v1               |
| `ReasoningFormatXaiResponsesV1`         | xai-responses-v1                        |
| `ReasoningFormatAnthropicClaudeV1`      | anthropic-claude-v1                     |
| `ReasoningFormatGoogleGeminiV1`         | google-gemini-v1                        |