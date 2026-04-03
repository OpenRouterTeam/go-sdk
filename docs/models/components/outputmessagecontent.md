# OutputMessageContent


## Supported Types

### ResponseOutputText

```go
outputMessageContent := components.CreateOutputMessageContentOutputText(components.ResponseOutputText{/* values here */})
```

### OpenAIResponsesRefusalContent

```go
outputMessageContent := components.CreateOutputMessageContentRefusal(components.OpenAIResponsesRefusalContent{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch outputMessageContent.Type {
	case components.OutputMessageContentTypeOutputText:
		// outputMessageContent.ResponseOutputText is populated
	case components.OutputMessageContentTypeRefusal:
		// outputMessageContent.OpenAIResponsesRefusalContent is populated
}
```
