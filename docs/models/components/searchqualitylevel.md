# SearchQualityLevel

How much context to retrieve per result. Defaults to medium (15000 chars). Applies to Exa and Parallel engines; ignored with native provider search and Firecrawl.

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