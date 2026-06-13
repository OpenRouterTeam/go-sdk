# GranularityName

Granularity identifier

## Example Usage

```go
import (
	"github.com/OpenRouterTeam/go-sdk/models/operations"
)

value := operations.GranularityNameMinute

// Open enum: custom values can be created with a direct type cast
custom := operations.GranularityName("custom_value")
```


## Values

| Name                    | Value                   |
| ----------------------- | ----------------------- |
| `GranularityNameMinute` | minute                  |
| `GranularityNameHour`   | hour                    |
| `GranularityNameDay`    | day                     |
| `GranularityNameWeek`   | week                    |
| `GranularityNameMonth`  | month                   |