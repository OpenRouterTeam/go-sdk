# VideoGenerationResponseStatus

## Example Usage

```go
import (
	"github.com/OpenRouterTeam/go-sdk/models/components"
)

value := components.VideoGenerationResponseStatusPending

// Open enum: custom values can be created with a direct type cast
custom := components.VideoGenerationResponseStatus("custom_value")
```


## Values

| Name                                      | Value                                     |
| ----------------------------------------- | ----------------------------------------- |
| `VideoGenerationResponseStatusPending`    | pending                                   |
| `VideoGenerationResponseStatusInProgress` | in_progress                               |
| `VideoGenerationResponseStatusCompleted`  | completed                                 |
| `VideoGenerationResponseStatusFailed`     | failed                                    |
| `VideoGenerationResponseStatusCancelled`  | cancelled                                 |
| `VideoGenerationResponseStatusExpired`    | expired                                   |