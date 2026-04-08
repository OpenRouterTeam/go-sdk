# ChatRequestServiceTier

The service tier to use for processing this request.

## Example Usage

```go
import (
	"github.com/OpenRouterTeam/go-sdk/models/components"
)

value := components.ChatRequestServiceTierAuto

// Open enum: custom values can be created with a direct type cast
custom := components.ChatRequestServiceTier("custom_value")
```


## Values

| Name                             | Value                            |
| -------------------------------- | -------------------------------- |
| `ChatRequestServiceTierAuto`     | auto                             |
| `ChatRequestServiceTierDefault`  | default                          |
| `ChatRequestServiceTierFlex`     | flex                             |
| `ChatRequestServiceTierPriority` | priority                         |
| `ChatRequestServiceTierScale`    | scale                            |