# WebFetchEngineEnum

Which fetch engine to use. "auto" (default) uses native if the provider supports it, otherwise Exa. "native" forces the provider's built-in fetch. "exa" uses Exa Contents API (supports BYOK). "openrouter" uses direct HTTP fetch. "firecrawl" uses Firecrawl scrape (requires BYOK).

## Example Usage

```go
import (
	"github.com/OpenRouterTeam/go-sdk/models/components"
)

value := components.WebFetchEngineEnumAuto

// Open enum: custom values can be created with a direct type cast
custom := components.WebFetchEngineEnum("custom_value")
```


## Values

| Name                           | Value                          |
| ------------------------------ | ------------------------------ |
| `WebFetchEngineEnumAuto`       | auto                           |
| `WebFetchEngineEnumNative`     | native                         |
| `WebFetchEngineEnumOpenrouter` | openrouter                     |
| `WebFetchEngineEnumFirecrawl`  | firecrawl                      |
| `WebFetchEngineEnumExa`        | exa                            |