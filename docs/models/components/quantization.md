# Quantization

## Example Usage

```go
import (
	"github.com/OpenRouterTeam/go-sdk/models/components"
)

value := components.QuantizationInt4

// Open enum: custom values can be created with a direct type cast
custom := components.Quantization("custom_value")
```


## Values

| Name                  | Value                 |
| --------------------- | --------------------- |
| `QuantizationInt4`    | int4                  |
| `QuantizationInt8`    | int8                  |
| `QuantizationFp4`     | fp4                   |
| `QuantizationFp6`     | fp6                   |
| `QuantizationFp8`     | fp8                   |
| `QuantizationFp16`    | fp16                  |
| `QuantizationBf16`    | bf16                  |
| `QuantizationFp32`    | fp32                  |
| `QuantizationUnknown` | unknown               |