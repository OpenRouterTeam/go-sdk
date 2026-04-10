# OpenAIResponseInputMessageItemContent


## Supported Types

### InputAudio

```go
openAIResponseInputMessageItemContent := components.CreateOpenAIResponseInputMessageItemContentInputAudio(components.InputAudio{/* values here */})
```

### InputFile

```go
openAIResponseInputMessageItemContent := components.CreateOpenAIResponseInputMessageItemContentInputFile(components.InputFile{/* values here */})
```

### InputImage

```go
openAIResponseInputMessageItemContent := components.CreateOpenAIResponseInputMessageItemContentInputImage(components.InputImage{/* values here */})
```

### InputText

```go
openAIResponseInputMessageItemContent := components.CreateOpenAIResponseInputMessageItemContentInputText(components.InputText{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch openAIResponseInputMessageItemContent.Type {
	case components.OpenAIResponseInputMessageItemContentTypeInputAudio:
		// openAIResponseInputMessageItemContent.InputAudio is populated
	case components.OpenAIResponseInputMessageItemContentTypeInputFile:
		// openAIResponseInputMessageItemContent.InputFile is populated
	case components.OpenAIResponseInputMessageItemContentTypeInputImage:
		// openAIResponseInputMessageItemContent.InputImage is populated
	case components.OpenAIResponseInputMessageItemContentTypeInputText:
		// openAIResponseInputMessageItemContent.InputText is populated
	default:
		// Unknown type - use openAIResponseInputMessageItemContent.GetUnknownRaw() for raw JSON
}
```
