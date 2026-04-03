# UsageLimitType

Optional credit limit reset interval. When set, the credit limit resets on this interval.

## Example Usage

```go
import (
	"github.com/OpenRouterTeam/go-sdk/models/operations"
)

value := operations.UsageLimitTypeDaily

// Open enum: custom values can be created with a direct type cast
custom := operations.UsageLimitType("custom_value")
```


## Values

| Name                    | Value                   |
| ----------------------- | ----------------------- |
| `UsageLimitTypeDaily`   | daily                   |
| `UsageLimitTypeWeekly`  | weekly                  |
| `UsageLimitTypeMonthly` | monthly                 |