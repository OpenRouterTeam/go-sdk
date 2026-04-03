# Preview20250311WebSearchServerToolEngine

Which search engine to use. "auto" (default) uses native if the provider supports it, otherwise Exa. "native" forces the provider's built-in search. "exa" forces the Exa search API. "firecrawl" uses Firecrawl (requires BYOK). "parallel" uses the Parallel search API.

## Example Usage

```go
import (
	"github.com/OpenRouterTeam/go-sdk/models/components"
)

value := components.Preview20250311WebSearchServerToolEngineAuto

// Open enum: custom values can be created with a direct type cast
custom := components.Preview20250311WebSearchServerToolEngine("custom_value")
```


## Values

| Name                                                | Value                                               |
| --------------------------------------------------- | --------------------------------------------------- |
| `Preview20250311WebSearchServerToolEngineAuto`      | auto                                                |
| `Preview20250311WebSearchServerToolEngineNative`    | native                                              |
| `Preview20250311WebSearchServerToolEngineExa`       | exa                                                 |
| `Preview20250311WebSearchServerToolEngineFirecrawl` | firecrawl                                           |
| `Preview20250311WebSearchServerToolEngineParallel`  | parallel                                            |