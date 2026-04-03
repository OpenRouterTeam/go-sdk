# ReasoningDetailSummaryFormat

## Example Usage

```go
import (
	"github.com/OpenRouterTeam/go-sdk/models/components"
)

value := components.ReasoningDetailSummaryFormatUnknown

// Open enum: custom values can be created with a direct type cast
custom := components.ReasoningDetailSummaryFormat("custom_value")
```


## Values

| Name                                                 | Value                                                |
| ---------------------------------------------------- | ---------------------------------------------------- |
| `ReasoningDetailSummaryFormatUnknown`                | unknown                                              |
| `ReasoningDetailSummaryFormatOpenaiResponsesV1`      | openai-responses-v1                                  |
| `ReasoningDetailSummaryFormatAzureOpenaiResponsesV1` | azure-openai-responses-v1                            |
| `ReasoningDetailSummaryFormatXaiResponsesV1`         | xai-responses-v1                                     |
| `ReasoningDetailSummaryFormatAnthropicClaudeV1`      | anthropic-claude-v1                                  |
| `ReasoningDetailSummaryFormatGoogleGeminiV1`         | google-gemini-v1                                     |