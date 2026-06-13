# DisplayFormat

How this metric value should be formatted for display (e.g. percent → multiply by 100 and append %, currency → prefix with $)

## Example Usage

```go
import (
	"github.com/OpenRouterTeam/go-sdk/models/operations"
)

value := operations.DisplayFormatNumber

// Open enum: custom values can be created with a direct type cast
custom := operations.DisplayFormat("custom_value")
```


## Values

| Name                      | Value                     |
| ------------------------- | ------------------------- |
| `DisplayFormatNumber`     | number                    |
| `DisplayFormatCurrency`   | currency                  |
| `DisplayFormatPercent`    | percent                   |
| `DisplayFormatLatency`    | latency                   |
| `DisplayFormatThroughput` | throughput                |