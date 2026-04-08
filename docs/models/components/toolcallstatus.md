# ToolCallStatus

## Example Usage

```go
import (
	"github.com/OpenRouterTeam/go-sdk/models/components"
)

value := components.ToolCallStatusInProgress

// Open enum: custom values can be created with a direct type cast
custom := components.ToolCallStatus("custom_value")
```


## Values

| Name                       | Value                      |
| -------------------------- | -------------------------- |
| `ToolCallStatusInProgress` | in_progress                |
| `ToolCallStatusCompleted`  | completed                  |
| `ToolCallStatusIncomplete` | incomplete                 |