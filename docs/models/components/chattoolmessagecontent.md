# ChatToolMessageContent

Tool response content


## Supported Types

### 

```go
chatToolMessageContent := components.CreateChatToolMessageContentStr(string{/* values here */})
```

### 

```go
chatToolMessageContent := components.CreateChatToolMessageContentArrayOfChatContentItems([]components.ChatContentItems{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch chatToolMessageContent.Type {
	case components.ChatToolMessageContentTypeStr:
		// chatToolMessageContent.Str is populated
	case components.ChatToolMessageContentTypeArrayOfChatContentItems:
		// chatToolMessageContent.ArrayOfChatContentItems is populated
}
```
