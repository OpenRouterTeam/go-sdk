# CreateGuardrailResetIntervalRequest

Interval at which the limit resets (daily, weekly, monthly)

## Example Usage

```go
import (
	"github.com/OpenRouterTeam/go-sdk/models/operations"
)

value := operations.CreateGuardrailResetIntervalRequestDaily

// Open enum: custom values can be created with a direct type cast
custom := operations.CreateGuardrailResetIntervalRequest("custom_value")
```


## Values

| Name                                         | Value                                        |
| -------------------------------------------- | -------------------------------------------- |
| `CreateGuardrailResetIntervalRequestDaily`   | daily                                        |
| `CreateGuardrailResetIntervalRequestWeekly`  | weekly                                       |
| `CreateGuardrailResetIntervalRequestMonthly` | monthly                                      |