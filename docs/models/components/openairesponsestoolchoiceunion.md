# OpenAIResponsesToolChoiceUnion


## Supported Types

### OpenAIResponsesToolChoiceAuto

```go
openAIResponsesToolChoiceUnion := components.CreateOpenAIResponsesToolChoiceUnionOpenAIResponsesToolChoiceAuto(components.OpenAIResponsesToolChoiceAuto{/* values here */})
```

### OpenAIResponsesToolChoiceNone

```go
openAIResponsesToolChoiceUnion := components.CreateOpenAIResponsesToolChoiceUnionOpenAIResponsesToolChoiceNone(components.OpenAIResponsesToolChoiceNone{/* values here */})
```

### OpenAIResponsesToolChoiceRequired

```go
openAIResponsesToolChoiceUnion := components.CreateOpenAIResponsesToolChoiceUnionOpenAIResponsesToolChoiceRequired(components.OpenAIResponsesToolChoiceRequired{/* values here */})
```

### OpenAIResponsesToolChoiceFunction

```go
openAIResponsesToolChoiceUnion := components.CreateOpenAIResponsesToolChoiceUnionOpenAIResponsesToolChoiceFunction(components.OpenAIResponsesToolChoiceFunction{/* values here */})
```

### OpenAIResponsesToolChoice

```go
openAIResponsesToolChoiceUnion := components.CreateOpenAIResponsesToolChoiceUnionOpenAIResponsesToolChoice(components.OpenAIResponsesToolChoice{/* values here */})
```

### ToolChoiceAllowed

```go
openAIResponsesToolChoiceUnion := components.CreateOpenAIResponsesToolChoiceUnionToolChoiceAllowed(components.ToolChoiceAllowed{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch openAIResponsesToolChoiceUnion.Type {
	case components.OpenAIResponsesToolChoiceUnionTypeOpenAIResponsesToolChoiceAuto:
		// openAIResponsesToolChoiceUnion.OpenAIResponsesToolChoiceAuto is populated
	case components.OpenAIResponsesToolChoiceUnionTypeOpenAIResponsesToolChoiceNone:
		// openAIResponsesToolChoiceUnion.OpenAIResponsesToolChoiceNone is populated
	case components.OpenAIResponsesToolChoiceUnionTypeOpenAIResponsesToolChoiceRequired:
		// openAIResponsesToolChoiceUnion.OpenAIResponsesToolChoiceRequired is populated
	case components.OpenAIResponsesToolChoiceUnionTypeOpenAIResponsesToolChoiceFunction:
		// openAIResponsesToolChoiceUnion.OpenAIResponsesToolChoiceFunction is populated
	case components.OpenAIResponsesToolChoiceUnionTypeOpenAIResponsesToolChoice:
		// openAIResponsesToolChoiceUnion.OpenAIResponsesToolChoice is populated
	case components.OpenAIResponsesToolChoiceUnionTypeToolChoiceAllowed:
		// openAIResponsesToolChoiceUnion.ToolChoiceAllowed is populated
}
```
