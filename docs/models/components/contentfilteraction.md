# ContentFilterAction

Action taken when the pattern matches

## Example Usage

```go
import (
	"github.com/OpenRouterTeam/go-sdk/models/components"
)

value := components.ContentFilterActionRedact

// Open enum: custom values can be created with a direct type cast
custom := components.ContentFilterAction("custom_value")
```


## Values

| Name                        | Value                       |
| --------------------------- | --------------------------- |
| `ContentFilterActionRedact` | redact                      |
| `ContentFilterActionBlock`  | block                       |