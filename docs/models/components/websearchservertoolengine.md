# WebSearchServerToolEngine

Which search engine to use. "auto" (default) uses native if the provider supports it, otherwise Exa. "native" forces the provider's built-in search. "exa" forces the Exa search API. "firecrawl" uses Firecrawl (requires BYOK). "parallel" uses the Parallel search API.

## Example Usage

```go
import (
	"github.com/OpenRouterTeam/go-sdk/models/components"
)

value := components.WebSearchServerToolEngineAuto

// Open enum: custom values can be created with a direct type cast
custom := components.WebSearchServerToolEngine("custom_value")
```


## Values

| Name                                 | Value                                |
| ------------------------------------ | ------------------------------------ |
| `WebSearchServerToolEngineAuto`      | auto                                 |
| `WebSearchServerToolEngineNative`    | native                               |
| `WebSearchServerToolEngineExa`       | exa                                  |
| `WebSearchServerToolEngineFirecrawl` | firecrawl                            |
| `WebSearchServerToolEngineParallel`  | parallel                             |