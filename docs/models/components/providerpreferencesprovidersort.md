# ProviderPreferencesProviderSort

The provider sorting strategy (price, throughput, latency)

## Example Usage

```go
import (
	"github.com/OpenRouterTeam/go-sdk/models/components"
)

value := components.ProviderPreferencesProviderSortPrice

// Open enum: custom values can be created with a direct type cast
custom := components.ProviderPreferencesProviderSort("custom_value")
```


## Values

| Name                                        | Value                                       |
| ------------------------------------------- | ------------------------------------------- |
| `ProviderPreferencesProviderSortPrice`      | price                                       |
| `ProviderPreferencesProviderSortThroughput` | throughput                                  |
| `ProviderPreferencesProviderSortLatency`    | latency                                     |
| `ProviderPreferencesProviderSortExacto`     | exacto                                      |