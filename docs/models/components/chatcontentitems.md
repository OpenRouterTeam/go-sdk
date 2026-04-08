# ChatContentItems

Content part for chat completion messages


## Supported Types

### ChatContentText

```go
chatContentItems := components.CreateChatContentItemsChatContentText(components.ChatContentText{/* values here */})
```

### ChatContentImage

```go
chatContentItems := components.CreateChatContentItemsChatContentImage(components.ChatContentImage{/* values here */})
```

### ChatContentAudio

```go
chatContentItems := components.CreateChatContentItemsChatContentAudio(components.ChatContentAudio{/* values here */})
```

### ChatContentItems1

```go
chatContentItems := components.CreateChatContentItemsChatContentItems1(components.ChatContentItems1{/* values here */})
```

### ChatContentFile

```go
chatContentItems := components.CreateChatContentItemsChatContentFile(components.ChatContentFile{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch chatContentItems.Type {
	case components.ChatContentItemsTypeChatContentText:
		// chatContentItems.ChatContentText is populated
	case components.ChatContentItemsTypeChatContentImage:
		// chatContentItems.ChatContentImage is populated
	case components.ChatContentItemsTypeChatContentAudio:
		// chatContentItems.ChatContentAudio is populated
	case components.ChatContentItemsTypeChatContentItems1:
		// chatContentItems.ChatContentItems1 is populated
	case components.ChatContentItemsTypeChatContentFile:
		// chatContentItems.ChatContentFile is populated
}
```
