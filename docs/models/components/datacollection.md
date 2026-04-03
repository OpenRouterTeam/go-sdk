# DataCollection

Data collection setting. If no available model provider meets the requirement, your request will return an error.
- allow: (default) allow providers which store user data non-transiently and may train on it

- deny: use only providers which do not collect user data.

## Example Usage

```go
import (
	"github.com/OpenRouterTeam/go-sdk/models/components"
)

value := components.DataCollectionDeny

// Open enum: custom values can be created with a direct type cast
custom := components.DataCollection("custom_value")
```


## Values

| Name                  | Value                 |
| --------------------- | --------------------- |
| `DataCollectionDeny`  | deny                  |
| `DataCollectionAllow` | allow                 |