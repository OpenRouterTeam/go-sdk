# UpdateGuardrailResetIntervalRequest

Interval at which the limit resets (daily, weekly, monthly)

## Example Usage

```go
import (
	"github.com/OpenRouterTeam/go-sdk/models/operations"
)

value := operations.UpdateGuardrailResetIntervalRequestDaily

// Open enum: custom values can be created with a direct type cast
custom := operations.UpdateGuardrailResetIntervalRequest("custom_value")
```


## Values

| Name                                         | Value                                        |
| -------------------------------------------- | -------------------------------------------- |
| `UpdateGuardrailResetIntervalRequestDaily`   | daily                                        |
| `UpdateGuardrailResetIntervalRequestWeekly`  | weekly                                       |
| `UpdateGuardrailResetIntervalRequestMonthly` | monthly                                      |