# SearchQualityLevel

How much context to retrieve per result. Defaults to medium (15000 chars). Only applies when using the Exa engine; ignored with native provider search.

## Example Usage

```go
import (
	"github.com/OpenRouterTeam/go-sdk/models/components"
)

value := components.SearchQualityLevelLow

// Open enum: custom values can be created with a direct type cast
custom := components.SearchQualityLevel("custom_value")
```


## Values

| Name                       | Value                      |
| -------------------------- | -------------------------- |
| `SearchQualityLevelLow`    | low                        |
| `SearchQualityLevelMedium` | medium                     |
| `SearchQualityLevelHigh`   | high                       |