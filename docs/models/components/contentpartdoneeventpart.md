# ContentPartDoneEventPart


## Supported Types

### ResponseOutputText

```go
contentPartDoneEventPart := components.CreateContentPartDoneEventPartOutputText(components.ResponseOutputText{/* values here */})
```

### ReasoningTextContent

```go
contentPartDoneEventPart := components.CreateContentPartDoneEventPartReasoningText(components.ReasoningTextContent{/* values here */})
```

### OpenAIResponsesRefusalContent

```go
contentPartDoneEventPart := components.CreateContentPartDoneEventPartRefusal(components.OpenAIResponsesRefusalContent{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch contentPartDoneEventPart.Type {
	case components.ContentPartDoneEventPartTypeOutputText:
		// contentPartDoneEventPart.ResponseOutputText is populated
	case components.ContentPartDoneEventPartTypeReasoningText:
		// contentPartDoneEventPart.ReasoningTextContent is populated
	case components.ContentPartDoneEventPartTypeRefusal:
		// contentPartDoneEventPart.OpenAIResponsesRefusalContent is populated
	default:
		// Unknown type - use contentPartDoneEventPart.GetUnknownRaw() for raw JSON
}
```
