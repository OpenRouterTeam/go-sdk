# UpdateGuardrailResetIntervalResponse

Interval at which the limit resets (daily, weekly, monthly)

## Example Usage

```go
import (
	"github.com/OpenRouterTeam/go-sdk/models/operations"
)

value := operations.UpdateGuardrailResetIntervalResponseDaily

// Open enum: custom values can be created with a direct type cast
custom := operations.UpdateGuardrailResetIntervalResponse("custom_value")
```


## Values

| Name                                          | Value                                         |
| --------------------------------------------- | --------------------------------------------- |
| `UpdateGuardrailResetIntervalResponseDaily`   | daily                                         |
| `UpdateGuardrailResetIntervalResponseWeekly`  | weekly                                        |
| `UpdateGuardrailResetIntervalResponseMonthly` | monthly                                       |