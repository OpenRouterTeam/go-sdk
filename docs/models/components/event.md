# Event

## Example Usage

```go
import (
	"github.com/OpenRouterTeam/go-sdk/models/components"
)

value := components.EventAdapterRequest

// Open enum: custom values can be created with a direct type cast
custom := components.Event("custom_value")
```


## Values

| Name                           | Value                          |
| ------------------------------ | ------------------------------ |
| `EventAdapterRequest`          | adapter_request                |
| `EventUpstreamHeadersReceived` | upstream_headers_received      |
| `EventFirstTokenReceived`      | first_token_received           |
| `EventUpstreamBodyEnded`       | upstream_body_ended            |