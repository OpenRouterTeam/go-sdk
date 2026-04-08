# APIType

Type of API used for the generation

## Example Usage

```go
import (
	"github.com/OpenRouterTeam/go-sdk/models/operations"
)

value := operations.APITypeCompletions

// Open enum: custom values can be created with a direct type cast
custom := operations.APIType("custom_value")
```


## Values

| Name                 | Value                |
| -------------------- | -------------------- |
| `APITypeCompletions` | completions          |
| `APITypeEmbeddings`  | embeddings           |
| `APITypeRerank`      | rerank               |
| `APITypeVideo`       | video                |