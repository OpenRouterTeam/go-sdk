# OpenAIResponseInputMessageItemContent


## Supported Types

### InputText

```go
openAIResponseInputMessageItemContent := components.CreateOpenAIResponseInputMessageItemContentInputText(components.InputText{/* values here */})
```

### InputImage

```go
openAIResponseInputMessageItemContent := components.CreateOpenAIResponseInputMessageItemContentInputImage(components.InputImage{/* values here */})
```

### InputFile

```go
openAIResponseInputMessageItemContent := components.CreateOpenAIResponseInputMessageItemContentInputFile(components.InputFile{/* values here */})
```

### InputAudio

```go
openAIResponseInputMessageItemContent := components.CreateOpenAIResponseInputMessageItemContentInputAudio(components.InputAudio{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch openAIResponseInputMessageItemContent.Type {
	case components.OpenAIResponseInputMessageItemContentTypeInputText:
		// openAIResponseInputMessageItemContent.InputText is populated
	case components.OpenAIResponseInputMessageItemContentTypeInputImage:
		// openAIResponseInputMessageItemContent.InputImage is populated
	case components.OpenAIResponseInputMessageItemContentTypeInputFile:
		// openAIResponseInputMessageItemContent.InputFile is populated
	case components.OpenAIResponseInputMessageItemContentTypeInputAudio:
		// openAIResponseInputMessageItemContent.InputAudio is populated
}
```
