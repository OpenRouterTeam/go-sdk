# ListGuardrailsResetInterval

Interval at which the limit resets (daily, weekly, monthly)

## Example Usage

```go
import (
	"github.com/OpenRouterTeam/go-sdk/models/operations"
)

value := operations.ListGuardrailsResetIntervalDaily

// Open enum: custom values can be created with a direct type cast
custom := operations.ListGuardrailsResetInterval("custom_value")
```


## Values

| Name                                 | Value                                |
| ------------------------------------ | ------------------------------------ |
| `ListGuardrailsResetIntervalDaily`   | daily                                |
| `ListGuardrailsResetIntervalWeekly`  | weekly                               |
| `ListGuardrailsResetIntervalMonthly` | monthly                              |