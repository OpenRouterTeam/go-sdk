# PublicEndpointQuantization

## Example Usage

```go
import (
	"github.com/OpenRouterTeam/go-sdk/models/components"
)

value := components.PublicEndpointQuantizationInt4

// Open enum: custom values can be created with a direct type cast
custom := components.PublicEndpointQuantization("custom_value")
```


## Values

| Name                                | Value                               |
| ----------------------------------- | ----------------------------------- |
| `PublicEndpointQuantizationInt4`    | int4                                |
| `PublicEndpointQuantizationInt8`    | int8                                |
| `PublicEndpointQuantizationFp4`     | fp4                                 |
| `PublicEndpointQuantizationFp6`     | fp6                                 |
| `PublicEndpointQuantizationFp8`     | fp8                                 |
| `PublicEndpointQuantizationFp16`    | fp16                                |
| `PublicEndpointQuantizationBf16`    | bf16                                |
| `PublicEndpointQuantizationFp32`    | fp32                                |
| `PublicEndpointQuantizationUnknown` | unknown                             |