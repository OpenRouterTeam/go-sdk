# ShellCallStatus

Status of a shell call or its output.

## Example Usage

```go
import (
	"github.com/OpenRouterTeam/go-sdk/models/components"
)

value := components.ShellCallStatusInProgress

// Open enum: custom values can be created with a direct type cast
custom := components.ShellCallStatus("custom_value")
```


## Values

| Name                        | Value                       |
| --------------------------- | --------------------------- |
| `ShellCallStatusInProgress` | in_progress                 |
| `ShellCallStatusCompleted`  | completed                   |
| `ShellCallStatusIncomplete` | incomplete                  |