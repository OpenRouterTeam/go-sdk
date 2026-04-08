# By

The provider sorting strategy (price, throughput, latency)

## Example Usage

```go
import (
	"github.com/OpenRouterTeam/go-sdk/models/components"
)

value := components.ByPrice

// Open enum: custom values can be created with a direct type cast
custom := components.By("custom_value")
```


## Values

| Name           | Value          |
| -------------- | -------------- |
| `ByPrice`      | price          |
| `ByThroughput` | throughput     |
| `ByLatency`    | latency        |
| `ByExacto`     | exacto         |