# GetAppRankingsSort

`popular` ranks apps by total token volume inside the date window. `trending` ranks apps by absolute excess token growth: window volume minus the average volume of the three equal-length periods immediately preceding the window. Apps with no excess growth are omitted from `trending` results.

## Example Usage

```go
import (
	"github.com/OpenRouterTeam/go-sdk/models/operations"
)

value := operations.GetAppRankingsSortPopular

// Open enum: custom values can be created with a direct type cast
custom := operations.GetAppRankingsSort("custom_value")
```


## Values

| Name                         | Value                        |
| ---------------------------- | ---------------------------- |
| `GetAppRankingsSortPopular`  | popular                      |
| `GetAppRankingsSortTrending` | trending                     |