# CreateObservabilityDestinationRequestType

The destination type. Only stable destination types are accepted.

## Example Usage

```go
import (
	"github.com/OpenRouterTeam/go-sdk/models/components"
)

value := components.CreateObservabilityDestinationRequestTypeArize

// Open enum: custom values can be created with a direct type cast
custom := components.CreateObservabilityDestinationRequestType("custom_value")
```


## Values

| Name                                                     | Value                                                    |
| -------------------------------------------------------- | -------------------------------------------------------- |
| `CreateObservabilityDestinationRequestTypeArize`         | arize                                                    |
| `CreateObservabilityDestinationRequestTypeBraintrust`    | braintrust                                               |
| `CreateObservabilityDestinationRequestTypeClickhouse`    | clickhouse                                               |
| `CreateObservabilityDestinationRequestTypeDatadog`       | datadog                                                  |
| `CreateObservabilityDestinationRequestTypeGrafana`       | grafana                                                  |
| `CreateObservabilityDestinationRequestTypeLangfuse`      | langfuse                                                 |
| `CreateObservabilityDestinationRequestTypeLangsmith`     | langsmith                                                |
| `CreateObservabilityDestinationRequestTypeNewrelic`      | newrelic                                                 |
| `CreateObservabilityDestinationRequestTypeOpik`          | opik                                                     |
| `CreateObservabilityDestinationRequestTypeOtelCollector` | otel-collector                                           |
| `CreateObservabilityDestinationRequestTypePosthog`       | posthog                                                  |
| `CreateObservabilityDestinationRequestTypeRamp`          | ramp                                                     |
| `CreateObservabilityDestinationRequestTypeS3`            | s3                                                       |
| `CreateObservabilityDestinationRequestTypeSentry`        | sentry                                                   |
| `CreateObservabilityDestinationRequestTypeSnowflake`     | snowflake                                                |
| `CreateObservabilityDestinationRequestTypeWeave`         | weave                                                    |
| `CreateObservabilityDestinationRequestTypeWebhook`       | webhook                                                  |