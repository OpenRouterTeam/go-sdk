# OperatorName

Operator identifier used in filter definitions

## Example Usage

```go
import (
	"github.com/OpenRouterTeam/go-sdk/models/operations"
)

value := operations.OperatorNameEq

// Open enum: custom values can be created with a direct type cast
custom := operations.OperatorName("custom_value")
```


## Values

| Name                | Value               |
| ------------------- | ------------------- |
| `OperatorNameEq`    | eq                  |
| `OperatorNameNeq`   | neq                 |
| `OperatorNameIn`    | in                  |
| `OperatorNameNotIn` | not_in              |
| `OperatorNameGt`    | gt                  |
| `OperatorNameGte`   | gte                 |
| `OperatorNameLt`    | lt                  |
| `OperatorNameLte`   | lte                 |