# ContentPartAddedEventPart


## Supported Types

### ResponseOutputText

```go
contentPartAddedEventPart := components.CreateContentPartAddedEventPartOutputText(components.ResponseOutputText{/* values here */})
```

### ReasoningTextContent

```go
contentPartAddedEventPart := components.CreateContentPartAddedEventPartReasoningText(components.ReasoningTextContent{/* values here */})
```

### OpenAIResponsesRefusalContent

```go
contentPartAddedEventPart := components.CreateContentPartAddedEventPartRefusal(components.OpenAIResponsesRefusalContent{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch contentPartAddedEventPart.Type {
	case components.ContentPartAddedEventPartTypeOutputText:
		// contentPartAddedEventPart.ResponseOutputText is populated
	case components.ContentPartAddedEventPartTypeReasoningText:
		// contentPartAddedEventPart.ReasoningTextContent is populated
	case components.ContentPartAddedEventPartTypeRefusal:
		// contentPartAddedEventPart.OpenAIResponsesRefusalContent is populated
	default:
		// Unknown type - use contentPartAddedEventPart.GetUnknownRaw() for raw JSON
}
```
