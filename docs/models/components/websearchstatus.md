# WebSearchStatus

## Example Usage

```go
import (
	"github.com/OpenRouterTeam/go-sdk/models/components"
)

value := components.WebSearchStatusCompleted

// Open enum: custom values can be created with a direct type cast
custom := components.WebSearchStatus("custom_value")
```


## Values

| Name                        | Value                       |
| --------------------------- | --------------------------- |
| `WebSearchStatusCompleted`  | completed                   |
| `WebSearchStatusSearching`  | searching                   |
| `WebSearchStatusInProgress` | in_progress                 |
| `WebSearchStatusFailed`     | failed                      |