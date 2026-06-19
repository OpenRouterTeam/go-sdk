# DefaultEffort

Default reasoning effort when the client enables reasoning without specifying effort. Maps to `reasoning.effort` in chat requests. When `"none"`, prefer omitting effort unless the user explicitly disables reasoning.

## Example Usage

```go
import (
	"github.com/OpenRouterTeam/go-sdk/models/components"
)

value := components.DefaultEffortXhigh

// Open enum: custom values can be created with a direct type cast
custom := components.DefaultEffort("custom_value")
```


## Values

| Name                   | Value                  |
| ---------------------- | ---------------------- |
| `DefaultEffortXhigh`   | xhigh                  |
| `DefaultEffortHigh`    | high                   |
| `DefaultEffortMedium`  | medium                 |
| `DefaultEffortLow`     | low                    |
| `DefaultEffortMinimal` | minimal                |
| `DefaultEffortNone`    | none                   |