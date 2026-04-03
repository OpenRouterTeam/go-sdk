# GetGuardrailResetInterval

Interval at which the limit resets (daily, weekly, monthly)

## Example Usage

```go
import (
	"github.com/OpenRouterTeam/go-sdk/models/operations"
)

value := operations.GetGuardrailResetIntervalDaily

// Open enum: custom values can be created with a direct type cast
custom := operations.GetGuardrailResetInterval("custom_value")
```


## Values

| Name                               | Value                              |
| ---------------------------------- | ---------------------------------- |
| `GetGuardrailResetIntervalDaily`   | daily                              |
| `GetGuardrailResetIntervalWeekly`  | weekly                             |
| `GetGuardrailResetIntervalMonthly` | monthly                            |