# ShellServerToolEngine

Which shell engine to use. "openrouter" runs commands server-side in the OpenRouter sandbox. "auto" (default) keeps the provider's native hosted shell when available (OpenAI); on other providers the call is routed to the OpenRouter sandbox.

## Example Usage

```go
import (
	"github.com/OpenRouterTeam/go-sdk/models/components"
)

value := components.ShellServerToolEngineAuto

// Open enum: custom values can be created with a direct type cast
custom := components.ShellServerToolEngine("custom_value")
```


## Values

| Name                              | Value                             |
| --------------------------------- | --------------------------------- |
| `ShellServerToolEngineAuto`       | auto                              |
| `ShellServerToolEngineOpenrouter` | openrouter                        |