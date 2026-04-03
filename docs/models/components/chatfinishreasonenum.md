# ChatFinishReasonEnum

## Example Usage

```go
import (
	"github.com/OpenRouterTeam/go-sdk/models/components"
)

value := components.ChatFinishReasonEnumToolCalls

// Open enum: custom values can be created with a direct type cast
custom := components.ChatFinishReasonEnum("custom_value")
```


## Values

| Name                                | Value                               |
| ----------------------------------- | ----------------------------------- |
| `ChatFinishReasonEnumToolCalls`     | tool_calls                          |
| `ChatFinishReasonEnumStop`          | stop                                |
| `ChatFinishReasonEnumLength`        | length                              |
| `ChatFinishReasonEnumContentFilter` | content_filter                      |
| `ChatFinishReasonEnumError`         | error                               |