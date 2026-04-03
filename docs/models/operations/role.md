# Role

Role of the member in the organization

## Example Usage

```go
import (
	"github.com/OpenRouterTeam/go-sdk/models/operations"
)

value := operations.RoleOrgAdmin

// Open enum: custom values can be created with a direct type cast
custom := operations.Role("custom_value")
```


## Values

| Name            | Value           |
| --------------- | --------------- |
| `RoleOrgAdmin`  | org:admin       |
| `RoleOrgMember` | org:member      |