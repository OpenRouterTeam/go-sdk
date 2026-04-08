# OutputMessageItemContent


## Supported Types

### ResponseOutputText

```go
outputMessageItemContent := components.CreateOutputMessageItemContentOutputText(components.ResponseOutputText{/* values here */})
```

### OpenAIResponsesRefusalContent

```go
outputMessageItemContent := components.CreateOutputMessageItemContentRefusal(components.OpenAIResponsesRefusalContent{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch outputMessageItemContent.Type {
	case components.OutputMessageItemContentTypeOutputText:
		// outputMessageItemContent.ResponseOutputText is populated
	case components.OutputMessageItemContentTypeRefusal:
		// outputMessageItemContent.OpenAIResponsesRefusalContent is populated
}
```
