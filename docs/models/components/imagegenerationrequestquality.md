# ImageGenerationRequestQuality

Rendering quality. Providers without a quality knob ignore this.

## Example Usage

```go
import (
	"github.com/OpenRouterTeam/go-sdk/models/components"
)

value := components.ImageGenerationRequestQualityAuto

// Open enum: custom values can be created with a direct type cast
custom := components.ImageGenerationRequestQuality("custom_value")
```


## Values

| Name                                  | Value                                 |
| ------------------------------------- | ------------------------------------- |
| `ImageGenerationRequestQualityAuto`   | auto                                  |
| `ImageGenerationRequestQualityLow`    | low                                   |
| `ImageGenerationRequestQualityMedium` | medium                                |
| `ImageGenerationRequestQualityHigh`   | high                                  |