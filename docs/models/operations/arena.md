# Arena

Design Arena only: arena to query. Defaults to `models` when source is `design-arena`.

## Example Usage

```go
import (
	"github.com/OpenRouterTeam/go-sdk/models/operations"
)

value := operations.ArenaModels

// Open enum: custom values can be created with a direct type cast
custom := operations.Arena("custom_value")
```


## Values

| Name            | Value           |
| --------------- | --------------- |
| `ArenaModels`   | models          |
| `ArenaBuilders` | builders        |
| `ArenaAgents`   | agents          |