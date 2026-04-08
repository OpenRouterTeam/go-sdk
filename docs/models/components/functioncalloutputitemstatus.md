# FunctionCallOutputItemStatus

## Example Usage

```go
import (
	"github.com/OpenRouterTeam/go-sdk/models/components"
)

value := components.FunctionCallOutputItemStatusInProgress

// Open enum: custom values can be created with a direct type cast
custom := components.FunctionCallOutputItemStatus("custom_value")
```


## Values

| Name                                     | Value                                    |
| ---------------------------------------- | ---------------------------------------- |
| `FunctionCallOutputItemStatusInProgress` | in_progress                              |
| `FunctionCallOutputItemStatusCompleted`  | completed                                |
| `FunctionCallOutputItemStatusIncomplete` | incomplete                               |