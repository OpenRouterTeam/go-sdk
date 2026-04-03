# PreviewWebSearchServerToolEngine

Which search engine to use. "auto" (default) uses native if the provider supports it, otherwise Exa. "native" forces the provider's built-in search. "exa" forces the Exa search API. "firecrawl" uses Firecrawl (requires BYOK). "parallel" uses the Parallel search API.

## Example Usage

```go
import (
	"github.com/OpenRouterTeam/go-sdk/models/components"
)

value := components.PreviewWebSearchServerToolEngineAuto

// Open enum: custom values can be created with a direct type cast
custom := components.PreviewWebSearchServerToolEngine("custom_value")
```


## Values

| Name                                        | Value                                       |
| ------------------------------------------- | ------------------------------------------- |
| `PreviewWebSearchServerToolEngineAuto`      | auto                                        |
| `PreviewWebSearchServerToolEngineNative`    | native                                      |
| `PreviewWebSearchServerToolEngineExa`       | exa                                         |
| `PreviewWebSearchServerToolEngineFirecrawl` | firecrawl                                   |
| `PreviewWebSearchServerToolEngineParallel`  | parallel                                    |