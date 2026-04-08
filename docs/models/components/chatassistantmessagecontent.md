# ChatAssistantMessageContent

Assistant message content


## Supported Types

### 

```go
chatAssistantMessageContent := components.CreateChatAssistantMessageContentStr(string{/* values here */})
```

### 

```go
chatAssistantMessageContent := components.CreateChatAssistantMessageContentArrayOfChatContentItems([]components.ChatContentItems{/* values here */})
```

### 

```go
chatAssistantMessageContent := components.CreateChatAssistantMessageContentAny(any{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch chatAssistantMessageContent.Type {
	case components.ChatAssistantMessageContentTypeStr:
		// chatAssistantMessageContent.Str is populated
	case components.ChatAssistantMessageContentTypeArrayOfChatContentItems:
		// chatAssistantMessageContent.ArrayOfChatContentItems is populated
	case components.ChatAssistantMessageContentTypeAny:
		// chatAssistantMessageContent.Any is populated
}
```
