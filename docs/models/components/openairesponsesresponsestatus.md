# OpenAIResponsesResponseStatus

## Example Usage

```go
import (
	"github.com/OpenRouterTeam/go-sdk/models/components"
)

value := components.OpenAIResponsesResponseStatusCompleted

// Open enum: custom values can be created with a direct type cast
custom := components.OpenAIResponsesResponseStatus("custom_value")
```


## Values

| Name                                      | Value                                     |
| ----------------------------------------- | ----------------------------------------- |
| `OpenAIResponsesResponseStatusCompleted`  | completed                                 |
| `OpenAIResponsesResponseStatusIncomplete` | incomplete                                |
| `OpenAIResponsesResponseStatusInProgress` | in_progress                               |
| `OpenAIResponsesResponseStatusFailed`     | failed                                    |
| `OpenAIResponsesResponseStatusCancelled`  | cancelled                                 |
| `OpenAIResponsesResponseStatusQueued`     | queued                                    |