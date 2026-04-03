# ChatRequestPartition

Partitioning strategy for sorting: "model" (default) groups endpoints by model before sorting (fallback models remain fallbacks), "none" sorts all endpoints together regardless of model.

## Example Usage

```go
import (
	"github.com/OpenRouterTeam/go-sdk/models/components"
)

value := components.ChatRequestPartitionModel

// Open enum: custom values can be created with a direct type cast
custom := components.ChatRequestPartition("custom_value")
```


## Values

| Name                        | Value                       |
| --------------------------- | --------------------------- |
| `ChatRequestPartitionModel` | model                       |
| `ChatRequestPartitionNone`  | none                        |