# Quality

## Example Usage

```go
import (
	"github.com/OpenRouterTeam/go-sdk/models/components"
)

value := components.QualityLow

// Open enum: custom values can be created with a direct type cast
custom := components.Quality("custom_value")
```


## Values

| Name            | Value           |
| --------------- | --------------- |
| `QualityLow`    | low             |
| `QualityMedium` | medium          |
| `QualityHigh`   | high            |
| `QualityAuto`   | auto            |