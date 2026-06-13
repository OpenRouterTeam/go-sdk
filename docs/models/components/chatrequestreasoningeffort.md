# ChatRequestReasoningEffort

Shorthand for setting reasoning effort. Equivalent to setting reasoning.effort. Cannot be used simultaneously with reasoning.effort if they differ.

## Example Usage

```go
import (
	"github.com/OpenRouterTeam/go-sdk/models/components"
)

value := components.ChatRequestReasoningEffortXhigh

// Open enum: custom values can be created with a direct type cast
custom := components.ChatRequestReasoningEffort("custom_value")
```


## Values

| Name                                | Value                               |
| ----------------------------------- | ----------------------------------- |
| `ChatRequestReasoningEffortXhigh`   | xhigh                               |
| `ChatRequestReasoningEffortHigh`    | high                                |
| `ChatRequestReasoningEffortMedium`  | medium                              |
| `ChatRequestReasoningEffortLow`     | low                                 |
| `ChatRequestReasoningEffortMinimal` | minimal                             |
| `ChatRequestReasoningEffortNone`    | none                                |