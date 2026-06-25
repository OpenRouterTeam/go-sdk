# GetModelsSort

Sort the returned models server-side. Prefer this over fetching the full list and sorting client-side. Options: pricing-low-to-high, pricing-high-to-low (average prompt/completion price), context-high-to-low (context length), throughput-high-to-low, latency-low-to-high (recent median performance), most-popular, top-weekly (tokens processed in the last week), newest (creation date), intelligence-high-to-low (Artificial Analysis intelligence index), design-arena-elo-high-to-low (best Design Arena ELO across arenas). Models without a score for the chosen benchmark are placed last. When omitted, the existing default ordering is preserved.

## Example Usage

```go
import (
	"github.com/OpenRouterTeam/go-sdk/models/operations"
)

value := operations.GetModelsSortMostPopular

// Open enum: custom values can be created with a direct type cast
custom := operations.GetModelsSort("custom_value")
```


## Values

| Name                                   | Value                                  |
| -------------------------------------- | -------------------------------------- |
| `GetModelsSortMostPopular`             | most-popular                           |
| `GetModelsSortNewest`                  | newest                                 |
| `GetModelsSortTopWeekly`               | top-weekly                             |
| `GetModelsSortPricingLowToHigh`        | pricing-low-to-high                    |
| `GetModelsSortPricingHighToLow`        | pricing-high-to-low                    |
| `GetModelsSortContextHighToLow`        | context-high-to-low                    |
| `GetModelsSortThroughputHighToLow`     | throughput-high-to-low                 |
| `GetModelsSortLatencyLowToHigh`        | latency-low-to-high                    |
| `GetModelsSortIntelligenceHighToLow`   | intelligence-high-to-low               |
| `GetModelsSortDesignArenaEloHighToLow` | design-arena-elo-high-to-low           |