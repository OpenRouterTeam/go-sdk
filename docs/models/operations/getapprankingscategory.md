# GetAppRankingsCategory

Marketplace category group to filter by (e.g. `coding`). Only apps tagged with a subcategory inside this group are returned. Mutually combinable with `subcategory` — when both are supplied the `subcategory` must belong to the `category` group.

## Example Usage

```go
import (
	"github.com/OpenRouterTeam/go-sdk/models/operations"
)

value := operations.GetAppRankingsCategoryCoding

// Open enum: custom values can be created with a direct type cast
custom := operations.GetAppRankingsCategory("custom_value")
```


## Values

| Name                                  | Value                                 |
| ------------------------------------- | ------------------------------------- |
| `GetAppRankingsCategoryCoding`        | coding                                |
| `GetAppRankingsCategoryCreative`      | creative                              |
| `GetAppRankingsCategoryProductivity`  | productivity                          |
| `GetAppRankingsCategoryEntertainment` | entertainment                         |