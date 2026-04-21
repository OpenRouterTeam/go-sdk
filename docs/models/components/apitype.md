# APIType

Type of API used for the generation

## Example Usage

```go
import (
	"github.com/OpenRouterTeam/go-sdk/models/components"
)

value := components.APITypeCompletions

// Open enum: custom values can be created with a direct type cast
custom := components.APIType("custom_value")
```


## Values

| Name                 | Value                |
| -------------------- | -------------------- |
| `APITypeCompletions` | completions          |
| `APITypeEmbeddings`  | embeddings           |
| `APITypeRerank`      | rerank               |
| `APITypeTts`         | tts                  |
| `APITypeVideo`       | video                |