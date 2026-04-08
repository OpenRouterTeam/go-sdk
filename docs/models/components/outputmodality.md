# OutputModality

## Example Usage

```go
import (
	"github.com/OpenRouterTeam/go-sdk/models/components"
)

value := components.OutputModalityText

// Open enum: custom values can be created with a direct type cast
custom := components.OutputModality("custom_value")
```


## Values

| Name                       | Value                      |
| -------------------------- | -------------------------- |
| `OutputModalityText`       | text                       |
| `OutputModalityImage`      | image                      |
| `OutputModalityEmbeddings` | embeddings                 |
| `OutputModalityAudio`      | audio                      |
| `OutputModalityVideo`      | video                      |
| `OutputModalityRerank`     | rerank                     |