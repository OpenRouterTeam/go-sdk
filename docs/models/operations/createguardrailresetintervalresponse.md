# CreateGuardrailResetIntervalResponse

Interval at which the limit resets (daily, weekly, monthly)

## Example Usage

```go
import (
	"github.com/OpenRouterTeam/go-sdk/models/operations"
)

value := operations.CreateGuardrailResetIntervalResponseDaily

// Open enum: custom values can be created with a direct type cast
custom := operations.CreateGuardrailResetIntervalResponse("custom_value")
```


## Values

| Name                                          | Value                                         |
| --------------------------------------------- | --------------------------------------------- |
| `CreateGuardrailResetIntervalResponseDaily`   | daily                                         |
| `CreateGuardrailResetIntervalResponseWeekly`  | weekly                                        |
| `CreateGuardrailResetIntervalResponseMonthly` | monthly                                       |