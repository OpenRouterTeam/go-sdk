# RoutingStrategy

## Example Usage

```go
import (
	"github.com/OpenRouterTeam/go-sdk/models/components"
)

value := components.RoutingStrategyDirect

// Open enum: custom values can be created with a direct type cast
custom := components.RoutingStrategy("custom_value")
```


## Values

| Name                         | Value                        |
| ---------------------------- | ---------------------------- |
| `RoutingStrategyDirect`      | direct                       |
| `RoutingStrategyAuto`        | auto                         |
| `RoutingStrategyFree`        | free                         |
| `RoutingStrategyLatest`      | latest                       |
| `RoutingStrategyAlias`       | alias                        |
| `RoutingStrategyFallback`    | fallback                     |
| `RoutingStrategyPareto`      | pareto                       |
| `RoutingStrategyBodybuilder` | bodybuilder                  |
| `RoutingStrategyFusion`      | fusion                       |