# PresetStatus

The status of a preset.

## Example Usage

```go
import (
	"github.com/OpenRouterTeam/go-sdk/models/components"
)

value := components.PresetStatusActive

// Open enum: custom values can be created with a direct type cast
custom := components.PresetStatus("custom_value")
```


## Values

| Name                   | Value                  |
| ---------------------- | ---------------------- |
| `PresetStatusActive`   | active                 |
| `PresetStatusDisabled` | disabled               |
| `PresetStatusArchived` | archived               |