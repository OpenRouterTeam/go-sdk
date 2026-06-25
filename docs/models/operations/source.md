# Source

Benchmark source to query. Determines the shape of the returned items. When omitted, returns results from all sources.

## Example Usage

```go
import (
	"github.com/OpenRouterTeam/go-sdk/models/operations"
)

value := operations.SourceArtificialAnalysis

// Open enum: custom values can be created with a direct type cast
custom := operations.Source("custom_value")
```


## Values

| Name                       | Value                      |
| -------------------------- | -------------------------- |
| `SourceArtificialAnalysis` | artificial-analysis        |
| `SourceDesignArena`        | design-arena               |