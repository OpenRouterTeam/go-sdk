# ChatToolChoice

Tool choice configuration


## Supported Types

### ChatToolChoiceNone

```go
chatToolChoice := components.CreateChatToolChoiceChatToolChoiceNone(components.ChatToolChoiceNone{/* values here */})
```

### ChatToolChoiceAuto

```go
chatToolChoice := components.CreateChatToolChoiceChatToolChoiceAuto(components.ChatToolChoiceAuto{/* values here */})
```

### ChatToolChoiceRequired

```go
chatToolChoice := components.CreateChatToolChoiceChatToolChoiceRequired(components.ChatToolChoiceRequired{/* values here */})
```

### ChatNamedToolChoice

```go
chatToolChoice := components.CreateChatToolChoiceChatNamedToolChoice(components.ChatNamedToolChoice{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch chatToolChoice.Type {
	case components.ChatToolChoiceTypeChatToolChoiceNone:
		// chatToolChoice.ChatToolChoiceNone is populated
	case components.ChatToolChoiceTypeChatToolChoiceAuto:
		// chatToolChoice.ChatToolChoiceAuto is populated
	case components.ChatToolChoiceTypeChatToolChoiceRequired:
		// chatToolChoice.ChatToolChoiceRequired is populated
	case components.ChatToolChoiceTypeChatNamedToolChoice:
		// chatToolChoice.ChatNamedToolChoice is populated
}
```
