# ChatRequestBy

The provider sorting strategy (price, throughput, latency)

## Example Usage

```go
import (
	"github.com/OpenRouterTeam/go-sdk/models/components"
)

value := components.ChatRequestByPrice

// Open enum: custom values can be created with a direct type cast
custom := components.ChatRequestBy("custom_value")
```


## Values

| Name                      | Value                     |
| ------------------------- | ------------------------- |
| `ChatRequestByPrice`      | price                     |
| `ChatRequestByThroughput` | throughput                |
| `ChatRequestByLatency`    | latency                   |
| `ChatRequestByExacto`     | exacto                    |