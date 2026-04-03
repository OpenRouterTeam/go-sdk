# CreateKeysLimitReset

Type of limit reset for the API key (daily, weekly, monthly, or null for no reset). Resets happen automatically at midnight UTC, and weeks are Monday through Sunday.

## Example Usage

```go
import (
	"github.com/OpenRouterTeam/go-sdk/models/operations"
)

value := operations.CreateKeysLimitResetDaily

// Open enum: custom values can be created with a direct type cast
custom := operations.CreateKeysLimitReset("custom_value")
```


## Values

| Name                          | Value                         |
| ----------------------------- | ----------------------------- |
| `CreateKeysLimitResetDaily`   | daily                         |
| `CreateKeysLimitResetWeekly`  | weekly                        |
| `CreateKeysLimitResetMonthly` | monthly                       |