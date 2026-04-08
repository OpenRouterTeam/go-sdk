# GuardrailInterval

Interval at which the limit resets (daily, weekly, monthly)

## Example Usage

```go
import (
	"github.com/OpenRouterTeam/go-sdk/models/components"
)

value := components.GuardrailIntervalDaily

// Open enum: custom values can be created with a direct type cast
custom := components.GuardrailInterval("custom_value")
```


## Values

| Name                       | Value                      |
| -------------------------- | -------------------------- |
| `GuardrailIntervalDaily`   | daily                      |
| `GuardrailIntervalWeekly`  | weekly                     |
| `GuardrailIntervalMonthly` | monthly                    |