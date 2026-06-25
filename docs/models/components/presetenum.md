# PresetEnum

A curated OpenRouter fusion preset (slugs follow `<task>-<tier>`, e.g. `general-high`). Expands server-side into the preset's analysis_models panel and judge model, so callers never name individual models. Explicitly provided `analysis_models` / `model` take precedence.

## Example Usage

```go
import (
	"github.com/OpenRouterTeam/go-sdk/models/components"
)

value := components.PresetEnumGeneralHigh

// Open enum: custom values can be created with a direct type cast
custom := components.PresetEnum("custom_value")
```


## Values

| Name                      | Value                     |
| ------------------------- | ------------------------- |
| `PresetEnumGeneralHigh`   | general-high              |
| `PresetEnumGeneralBudget` | general-budget            |
| `PresetEnumGeneralFast`   | general-fast              |