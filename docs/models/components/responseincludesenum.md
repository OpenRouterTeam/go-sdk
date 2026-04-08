# ResponseIncludesEnum

## Example Usage

```go
import (
	"github.com/OpenRouterTeam/go-sdk/models/components"
)

value := components.ResponseIncludesEnumFileSearchCallResults

// Open enum: custom values can be created with a direct type cast
custom := components.ResponseIncludesEnum("custom_value")
```


## Values

| Name                                                   | Value                                                  |
| ------------------------------------------------------ | ------------------------------------------------------ |
| `ResponseIncludesEnumFileSearchCallResults`            | file_search_call.results                               |
| `ResponseIncludesEnumMessageInputImageImageURL`        | message.input_image.image_url                          |
| `ResponseIncludesEnumComputerCallOutputOutputImageURL` | computer_call_output.output.image_url                  |
| `ResponseIncludesEnumReasoningEncryptedContent`        | reasoning.encrypted_content                            |
| `ResponseIncludesEnumCodeInterpreterCallOutputs`       | code_interpreter_call.outputs                          |