# InputModality

## Example Usage

```go
import (
	"github.com/OpenRouterTeam/go-sdk/models/components"
)

value := components.InputModalityText

// Open enum: custom values can be created with a direct type cast
custom := components.InputModality("custom_value")
```


## Values

| Name                 | Value                |
| -------------------- | -------------------- |
| `InputModalityText`  | text                 |
| `InputModalityImage` | image                |
| `InputModalityFile`  | file                 |
| `InputModalityAudio` | audio                |
| `InputModalityVideo` | video                |