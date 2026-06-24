# ResetInterval

Interval at which spend resets. Null means a lifetime (one-time) budget.

## Example Usage

```go
import (
	"github.com/OpenRouterTeam/go-sdk/models/components"
)

value := components.ResetIntervalDaily

// Open enum: custom values can be created with a direct type cast
custom := components.ResetInterval("custom_value")
```


## Values

| Name                   | Value                  |
| ---------------------- | ---------------------- |
| `ResetIntervalDaily`   | daily                  |
| `ResetIntervalWeekly`  | weekly                 |
| `ResetIntervalMonthly` | monthly                |