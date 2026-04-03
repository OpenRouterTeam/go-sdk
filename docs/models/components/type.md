# Type


## Supported Types

### OpenAIResponsesToolChoiceTypeWebSearchPreview20250311

```go
type := components.CreateTypeOpenAIResponsesToolChoiceTypeWebSearchPreview20250311(components.OpenAIResponsesToolChoiceTypeWebSearchPreview20250311{/* values here */})
```

### OpenAIResponsesToolChoiceTypeWebSearchPreview

```go
type := components.CreateTypeOpenAIResponsesToolChoiceTypeWebSearchPreview(components.OpenAIResponsesToolChoiceTypeWebSearchPreview{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch type.Type {
	case components.TypeTypeOpenAIResponsesToolChoiceTypeWebSearchPreview20250311:
		// type.OpenAIResponsesToolChoiceTypeWebSearchPreview20250311 is populated
	case components.TypeTypeOpenAIResponsesToolChoiceTypeWebSearchPreview:
		// type.OpenAIResponsesToolChoiceTypeWebSearchPreview is populated
}
```
