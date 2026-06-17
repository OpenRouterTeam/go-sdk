# OpenAIResponseCustomToolCallOutputOutput1


## Supported Types

### InputFile

```go
openAIResponseCustomToolCallOutputOutput1 := components.CreateOpenAIResponseCustomToolCallOutputOutput1InputFile(components.InputFile{/* values here */})
```

### InputImage

```go
openAIResponseCustomToolCallOutputOutput1 := components.CreateOpenAIResponseCustomToolCallOutputOutput1InputImage(components.InputImage{/* values here */})
```

### InputText

```go
openAIResponseCustomToolCallOutputOutput1 := components.CreateOpenAIResponseCustomToolCallOutputOutput1InputText(components.InputText{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch openAIResponseCustomToolCallOutputOutput1.Type {
	case components.OpenAIResponseCustomToolCallOutputOutput1TypeInputFile:
		// openAIResponseCustomToolCallOutputOutput1.InputFile is populated
	case components.OpenAIResponseCustomToolCallOutputOutput1TypeInputImage:
		// openAIResponseCustomToolCallOutputOutput1.InputImage is populated
	case components.OpenAIResponseCustomToolCallOutputOutput1TypeInputText:
		// openAIResponseCustomToolCallOutputOutput1.InputText is populated
	default:
		// Unknown type - use openAIResponseCustomToolCallOutputOutput1.GetUnknownRaw() for raw JSON
}
```
