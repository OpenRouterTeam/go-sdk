# ImageGenerationRequestBackground

Background treatment. `transparent` requires an output_format that supports alpha (png or webp).

## Example Usage

```go
import (
	"github.com/OpenRouterTeam/go-sdk/models/components"
)

value := components.ImageGenerationRequestBackgroundAuto

// Open enum: custom values can be created with a direct type cast
custom := components.ImageGenerationRequestBackground("custom_value")
```


## Values

| Name                                          | Value                                         |
| --------------------------------------------- | --------------------------------------------- |
| `ImageGenerationRequestBackgroundAuto`        | auto                                          |
| `ImageGenerationRequestBackgroundTransparent` | transparent                                   |
| `ImageGenerationRequestBackgroundOpaque`      | opaque                                        |