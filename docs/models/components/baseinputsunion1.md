# BaseInputsUnion1


## Supported Types

### BaseInputsMessage

```go
baseInputsUnion1 := components.CreateBaseInputsUnion1BaseInputsMessage(components.BaseInputsMessage{/* values here */})
```

### OpenAIResponseInputMessageItem

```go
baseInputsUnion1 := components.CreateBaseInputsUnion1OpenAIResponseInputMessageItem(components.OpenAIResponseInputMessageItem{/* values here */})
```

### OpenAIResponseFunctionToolCallOutput

```go
baseInputsUnion1 := components.CreateBaseInputsUnion1OpenAIResponseFunctionToolCallOutput(components.OpenAIResponseFunctionToolCallOutput{/* values here */})
```

### OpenAIResponseFunctionToolCall

```go
baseInputsUnion1 := components.CreateBaseInputsUnion1OpenAIResponseFunctionToolCall(components.OpenAIResponseFunctionToolCall{/* values here */})
```

### OutputItemImageGenerationCall

```go
baseInputsUnion1 := components.CreateBaseInputsUnion1OutputItemImageGenerationCall(components.OutputItemImageGenerationCall{/* values here */})
```

### OutputMessage

```go
baseInputsUnion1 := components.CreateBaseInputsUnion1OutputMessage(components.OutputMessage{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch baseInputsUnion1.Type {
	case components.BaseInputsUnion1TypeBaseInputsMessage:
		// baseInputsUnion1.BaseInputsMessage is populated
	case components.BaseInputsUnion1TypeOpenAIResponseInputMessageItem:
		// baseInputsUnion1.OpenAIResponseInputMessageItem is populated
	case components.BaseInputsUnion1TypeOpenAIResponseFunctionToolCallOutput:
		// baseInputsUnion1.OpenAIResponseFunctionToolCallOutput is populated
	case components.BaseInputsUnion1TypeOpenAIResponseFunctionToolCall:
		// baseInputsUnion1.OpenAIResponseFunctionToolCall is populated
	case components.BaseInputsUnion1TypeOutputItemImageGenerationCall:
		// baseInputsUnion1.OutputItemImageGenerationCall is populated
	case components.BaseInputsUnion1TypeOutputMessage:
		// baseInputsUnion1.OutputMessage is populated
}
```
