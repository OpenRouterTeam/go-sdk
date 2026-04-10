# OpenAIResponseFunctionToolCallOutputOutput1


## Supported Types

### InputFile

```go
openAIResponseFunctionToolCallOutputOutput1 := components.CreateOpenAIResponseFunctionToolCallOutputOutput1InputFile(components.InputFile{/* values here */})
```

### InputImage

```go
openAIResponseFunctionToolCallOutputOutput1 := components.CreateOpenAIResponseFunctionToolCallOutputOutput1InputImage(components.InputImage{/* values here */})
```

### InputText

```go
openAIResponseFunctionToolCallOutputOutput1 := components.CreateOpenAIResponseFunctionToolCallOutputOutput1InputText(components.InputText{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch openAIResponseFunctionToolCallOutputOutput1.Type {
	case components.OpenAIResponseFunctionToolCallOutputOutput1TypeInputFile:
		// openAIResponseFunctionToolCallOutputOutput1.InputFile is populated
	case components.OpenAIResponseFunctionToolCallOutputOutput1TypeInputImage:
		// openAIResponseFunctionToolCallOutputOutput1.InputImage is populated
	case components.OpenAIResponseFunctionToolCallOutputOutput1TypeInputText:
		// openAIResponseFunctionToolCallOutputOutput1.InputText is populated
	default:
		// Unknown type - use openAIResponseFunctionToolCallOutputOutput1.GetUnknownRaw() for raw JSON
}
```
