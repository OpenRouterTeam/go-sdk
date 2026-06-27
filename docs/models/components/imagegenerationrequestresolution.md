# ImageGenerationRequestResolution

Normalized resolution tier of the generated image. Concrete pixel dimensions are derived per-provider.

## Example Usage

```go
import (
	"github.com/OpenRouterTeam/go-sdk/models/components"
)

value := components.ImageGenerationRequestResolutionFiveHundredAndTwelve

// Open enum: custom values can be created with a direct type cast
custom := components.ImageGenerationRequestResolution("custom_value")
```


## Values

| Name                                                   | Value                                                  |
| ------------------------------------------------------ | ------------------------------------------------------ |
| `ImageGenerationRequestResolutionFiveHundredAndTwelve` | 512                                                    |
| `ImageGenerationRequestResolutionOneK`                 | 1K                                                     |
| `ImageGenerationRequestResolutionTwoK`                 | 2K                                                     |
| `ImageGenerationRequestResolutionFourK`                | 4K                                                     |