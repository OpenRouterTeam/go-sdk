# LegacyWebSearchServerToolEngine

Which search engine to use. "auto" (default) uses native if the provider supports it, otherwise Exa. "native" forces the provider's built-in search. "exa" forces the Exa search API. "firecrawl" uses Firecrawl (requires BYOK). "parallel" uses the Parallel search API.

## Example Usage

```go
import (
	"github.com/OpenRouterTeam/go-sdk/models/components"
)

value := components.LegacyWebSearchServerToolEngineAuto

// Open enum: custom values can be created with a direct type cast
custom := components.LegacyWebSearchServerToolEngine("custom_value")
```


## Values

| Name                                       | Value                                      |
| ------------------------------------------ | ------------------------------------------ |
| `LegacyWebSearchServerToolEngineAuto`      | auto                                       |
| `LegacyWebSearchServerToolEngineNative`    | native                                     |
| `LegacyWebSearchServerToolEngineExa`       | exa                                        |
| `LegacyWebSearchServerToolEngineFirecrawl` | firecrawl                                  |
| `LegacyWebSearchServerToolEngineParallel`  | parallel                                   |