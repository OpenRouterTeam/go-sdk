# FrameType

Whether this image represents the first or last frame of the video

## Example Usage

```go
import (
	"github.com/OpenRouterTeam/go-sdk/models/components"
)

value := components.FrameTypeFirstFrame

// Open enum: custom values can be created with a direct type cast
custom := components.FrameType("custom_value")
```


## Values

| Name                  | Value                 |
| --------------------- | --------------------- |
| `FrameTypeFirstFrame` | first_frame           |
| `FrameTypeLastFrame`  | last_frame            |