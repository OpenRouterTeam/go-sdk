# DataRegion

The data region this generation was routed through. 'europe' for EU-routed requests, 'global' otherwise.

## Example Usage

```go
import (
	"github.com/OpenRouterTeam/go-sdk/models/components"
)

value := components.DataRegionGlobal

// Open enum: custom values can be created with a direct type cast
custom := components.DataRegion("custom_value")
```


## Values

| Name               | Value              |
| ------------------ | ------------------ |
| `DataRegionGlobal` | global             |
| `DataRegionEurope` | europe             |