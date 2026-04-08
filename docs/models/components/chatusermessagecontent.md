# ChatUserMessageContent

User message content


## Supported Types

### 

```go
chatUserMessageContent := components.CreateChatUserMessageContentStr(string{/* values here */})
```

### 

```go
chatUserMessageContent := components.CreateChatUserMessageContentArrayOfChatContentItems([]components.ChatContentItems{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch chatUserMessageContent.Type {
	case components.ChatUserMessageContentTypeStr:
		// chatUserMessageContent.Str is populated
	case components.ChatUserMessageContentTypeArrayOfChatContentItems:
		// chatUserMessageContent.ArrayOfChatContentItems is populated
}
```
