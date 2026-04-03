# WebSearchEngine

The search engine to use for web search.

## Example Usage

```go
import (
	"github.com/OpenRouterTeam/go-sdk/models/components"
)

value := components.WebSearchEngineNative

// Open enum: custom values can be created with a direct type cast
custom := components.WebSearchEngine("custom_value")
```


## Values

| Name                       | Value                      |
| -------------------------- | -------------------------- |
| `WebSearchEngineNative`    | native                     |
| `WebSearchEngineExa`       | exa                        |
| `WebSearchEngineFirecrawl` | firecrawl                  |
| `WebSearchEngineParallel`  | parallel                   |