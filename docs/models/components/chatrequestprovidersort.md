# ChatRequestProviderSort

The provider sorting strategy (price, throughput, latency)

## Example Usage

```go
import (
	"github.com/OpenRouterTeam/go-sdk/models/components"
)

value := components.ChatRequestProviderSortPrice

// Open enum: custom values can be created with a direct type cast
custom := components.ChatRequestProviderSort("custom_value")
```


## Values

| Name                                | Value                               |
| ----------------------------------- | ----------------------------------- |
| `ChatRequestProviderSortPrice`      | price                               |
| `ChatRequestProviderSortThroughput` | throughput                          |
| `ChatRequestProviderSortLatency`    | latency                             |
| `ChatRequestProviderSortExacto`     | exacto                              |