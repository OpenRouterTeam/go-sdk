# WebSearchEngineEnum

Which search engine to use. "auto" (default) uses native if the provider supports it, otherwise Exa. "native" forces the provider's built-in search. "exa" forces the Exa search API. "firecrawl" uses Firecrawl (requires BYOK). "parallel" uses the Parallel search API. "perplexity" uses the Perplexity Search API (raw ranked results).

## Example Usage

```go
import (
	"github.com/OpenRouterTeam/go-sdk/models/components"
)

value := components.WebSearchEngineEnumNative

// Open enum: custom values can be created with a direct type cast
custom := components.WebSearchEngineEnum("custom_value")
```


## Values

| Name                            | Value                           |
| ------------------------------- | ------------------------------- |
| `WebSearchEngineEnumNative`     | native                          |
| `WebSearchEngineEnumExa`        | exa                             |
| `WebSearchEngineEnumParallel`   | parallel                        |
| `WebSearchEngineEnumFirecrawl`  | firecrawl                       |
| `WebSearchEngineEnumPerplexity` | perplexity                      |
| `WebSearchEngineEnumAuto`       | auto                            |