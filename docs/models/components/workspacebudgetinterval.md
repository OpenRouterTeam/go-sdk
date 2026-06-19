# WorkspaceBudgetInterval

Budget reset interval. Use "lifetime" for a one-time budget that never resets.

## Example Usage

```go
import (
	"github.com/OpenRouterTeam/go-sdk/models/components"
)

value := components.WorkspaceBudgetIntervalDaily

// Open enum: custom values can be created with a direct type cast
custom := components.WorkspaceBudgetInterval("custom_value")
```


## Values

| Name                              | Value                             |
| --------------------------------- | --------------------------------- |
| `WorkspaceBudgetIntervalDaily`    | daily                             |
| `WorkspaceBudgetIntervalWeekly`   | weekly                            |
| `WorkspaceBudgetIntervalMonthly`  | monthly                           |
| `WorkspaceBudgetIntervalLifetime` | lifetime                          |