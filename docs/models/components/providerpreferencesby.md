# ProviderPreferencesBy

The provider sorting strategy (price, throughput, latency)

## Example Usage

```go
import (
	"github.com/OpenRouterTeam/go-sdk/models/components"
)

value := components.ProviderPreferencesByPrice

// Open enum: custom values can be created with a direct type cast
custom := components.ProviderPreferencesBy("custom_value")
```


## Values

| Name                              | Value                             |
| --------------------------------- | --------------------------------- |
| `ProviderPreferencesByPrice`      | price                             |
| `ProviderPreferencesByThroughput` | throughput                        |
| `ProviderPreferencesByLatency`    | latency                           |
| `ProviderPreferencesByExacto`     | exacto                            |