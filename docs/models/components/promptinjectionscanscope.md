# PromptInjectionScanScope

Which message roles to scan for prompt injection. Only applies to the regex-prompt-injection builtin. Defaults to all_messages.

## Example Usage

```go
import (
	"github.com/OpenRouterTeam/go-sdk/models/components"
)

value := components.PromptInjectionScanScopeUserOnly

// Open enum: custom values can be created with a direct type cast
custom := components.PromptInjectionScanScope("custom_value")
```


## Values

| Name                                  | Value                                 |
| ------------------------------------- | ------------------------------------- |
| `PromptInjectionScanScopeUserOnly`    | user_only                             |
| `PromptInjectionScanScopeAllMessages` | all_messages                          |