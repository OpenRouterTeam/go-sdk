# ValueType

Whether the operator expects a single value or an array

## Example Usage

```go
import (
	"github.com/OpenRouterTeam/go-sdk/models/operations"
)

value := operations.ValueTypeScalar

// Open enum: custom values can be created with a direct type cast
custom := operations.ValueType("custom_value")
```


## Values

| Name              | Value             |
| ----------------- | ----------------- |
| `ValueTypeScalar` | scalar            |
| `ValueTypeArray`  | array             |