# ProviderSort

The provider sorting strategy (price, throughput, latency)

## Example Usage

```go
import (
	"github.com/OpenRouterTeam/go-sdk/models/components"
)

value := components.ProviderSortPrice

// Open enum: custom values can be created with a direct type cast
custom := components.ProviderSort("custom_value")
```


## Values

| Name                     | Value                    |
| ------------------------ | ------------------------ |
| `ProviderSortPrice`      | price                    |
| `ProviderSortThroughput` | throughput               |
| `ProviderSortLatency`    | latency                  |
| `ProviderSortExacto`     | exacto                   |