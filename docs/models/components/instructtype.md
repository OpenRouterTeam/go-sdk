# InstructType

Instruction format type

## Example Usage

```go
import (
	"github.com/OpenRouterTeam/go-sdk/models/components"
)

value := components.InstructTypeNone

// Open enum: custom values can be created with a direct type cast
custom := components.InstructType("custom_value")
```


## Values

| Name                      | Value                     |
| ------------------------- | ------------------------- |
| `InstructTypeNone`        | none                      |
| `InstructTypeAiroboros`   | airoboros                 |
| `InstructTypeAlpaca`      | alpaca                    |
| `InstructTypeAlpacaModif` | alpaca-modif              |
| `InstructTypeChatml`      | chatml                    |
| `InstructTypeClaude`      | claude                    |
| `InstructTypeCodeLlama`   | code-llama                |
| `InstructTypeGemma`       | gemma                     |
| `InstructTypeLlama2`      | llama2                    |
| `InstructTypeLlama3`      | llama3                    |
| `InstructTypeMistral`     | mistral                   |
| `InstructTypeNemotron`    | nemotron                  |
| `InstructTypeNeural`      | neural                    |
| `InstructTypeOpenchat`    | openchat                  |
| `InstructTypePhi3`        | phi3                      |
| `InstructTypeRwkv`        | rwkv                      |
| `InstructTypeVicuna`      | vicuna                    |
| `InstructTypeZephyr`      | zephyr                    |
| `InstructTypeDeepseekR1`  | deepseek-r1               |
| `InstructTypeDeepseekV31` | deepseek-v3.1             |
| `InstructTypeQwq`         | qwq                       |
| `InstructTypeQwen3`       | qwen3                     |