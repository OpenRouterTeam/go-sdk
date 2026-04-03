# ImageGenerationStatus

## Example Usage

```go
import (
	"github.com/OpenRouterTeam/go-sdk/models/components"
)

value := components.ImageGenerationStatusInProgress

// Open enum: custom values can be created with a direct type cast
custom := components.ImageGenerationStatus("custom_value")
```


## Values

| Name                              | Value                             |
| --------------------------------- | --------------------------------- |
| `ImageGenerationStatusInProgress` | in_progress                       |
| `ImageGenerationStatusCompleted`  | completed                         |
| `ImageGenerationStatusGenerating` | generating                        |
| `ImageGenerationStatusFailed`     | failed                            |