# SearchQualityLevel

How much context to retrieve per result. Applies to Exa, Parallel, and Perplexity engines; ignored with native provider search and Firecrawl. For Exa, pins a fixed per-result character cap (low=5,000, medium=15,000, high=30,000); when omitted, Exa picks an adaptive size per query and document (typically ~2,000–4,000 characters per result). For Parallel, controls the total characters across all results; when omitted, Parallel uses its own default size. For Perplexity, maps directly to the Search API's native search_context_size parameter. Overridden by `max_characters` when both are set.

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