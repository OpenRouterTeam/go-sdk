# ChatWebSearchShorthandSearchContextSize

How much context to retrieve per result. Defaults to medium (15000 chars). Only applies when using the Exa engine; ignored with native provider search.

## Example Usage

```go
import (
	"github.com/OpenRouterTeam/go-sdk/models/components"
)

value := components.ChatWebSearchShorthandSearchContextSizeLow

// Open enum: custom values can be created with a direct type cast
custom := components.ChatWebSearchShorthandSearchContextSize("custom_value")
```


## Values

| Name                                            | Value                                           |
| ----------------------------------------------- | ----------------------------------------------- |
| `ChatWebSearchShorthandSearchContextSizeLow`    | low                                             |
| `ChatWebSearchShorthandSearchContextSizeMedium` | medium                                          |
| `ChatWebSearchShorthandSearchContextSizeHigh`   | high                                            |