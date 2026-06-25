# ImageGenerationRequestOutputFormat

Encoding of the returned image bytes.

## Example Usage

```go
import (
	"github.com/OpenRouterTeam/go-sdk/models/components"
)

value := components.ImageGenerationRequestOutputFormatPng

// Open enum: custom values can be created with a direct type cast
custom := components.ImageGenerationRequestOutputFormat("custom_value")
```


## Values

| Name                                     | Value                                    |
| ---------------------------------------- | ---------------------------------------- |
| `ImageGenerationRequestOutputFormatPng`  | png                                      |
| `ImageGenerationRequestOutputFormatJpeg` | jpeg                                     |
| `ImageGenerationRequestOutputFormatWebp` | webp                                     |