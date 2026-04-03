# ChatMessages

Chat completion message with role-based discrimination


## Supported Types

### ChatSystemMessage

```go
chatMessages := components.CreateChatMessagesSystem(components.ChatSystemMessage{/* values here */})
```

### ChatUserMessage

```go
chatMessages := components.CreateChatMessagesUser(components.ChatUserMessage{/* values here */})
```

### ChatDeveloperMessage

```go
chatMessages := components.CreateChatMessagesDeveloper(components.ChatDeveloperMessage{/* values here */})
```

### ChatAssistantMessage

```go
chatMessages := components.CreateChatMessagesAssistant(components.ChatAssistantMessage{/* values here */})
```

### ChatToolMessage

```go
chatMessages := components.CreateChatMessagesTool(components.ChatToolMessage{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch chatMessages.Type {
	case components.ChatMessagesTypeSystem:
		// chatMessages.ChatSystemMessage is populated
	case components.ChatMessagesTypeUser:
		// chatMessages.ChatUserMessage is populated
	case components.ChatMessagesTypeDeveloper:
		// chatMessages.ChatDeveloperMessage is populated
	case components.ChatMessagesTypeAssistant:
		// chatMessages.ChatAssistantMessage is populated
	case components.ChatMessagesTypeTool:
		// chatMessages.ChatToolMessage is populated
}
```
