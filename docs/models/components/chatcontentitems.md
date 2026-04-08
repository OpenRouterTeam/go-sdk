# ChatContentItems

Content part for chat completion messages


## Supported Types

### ChatContentText

```go
chatContentItems := components.CreateChatContentItemsText(components.ChatContentText{/* values here */})
```

### ChatContentImage

```go
chatContentItems := components.CreateChatContentItemsImageURL(components.ChatContentImage{/* values here */})
```

### ChatContentAudio

```go
chatContentItems := components.CreateChatContentItemsInputAudio(components.ChatContentAudio{/* values here */})
```

### LegacyChatContentVideo

```go
chatContentItems := components.CreateChatContentItemsInputVideo(components.LegacyChatContentVideo{/* values here */})
```

### ChatContentVideo

```go
chatContentItems := components.CreateChatContentItemsVideoURL(components.ChatContentVideo{/* values here */})
```

### ChatContentFile

```go
chatContentItems := components.CreateChatContentItemsFile(components.ChatContentFile{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch chatContentItems.Type {
	case components.ChatContentItemsTypeText:
		// chatContentItems.ChatContentText is populated
	case components.ChatContentItemsTypeImageURL:
		// chatContentItems.ChatContentImage is populated
	case components.ChatContentItemsTypeInputAudio:
		// chatContentItems.ChatContentAudio is populated
	case components.ChatContentItemsTypeInputVideo:
		// chatContentItems.LegacyChatContentVideo is populated
	case components.ChatContentItemsTypeVideoURL:
		// chatContentItems.ChatContentVideo is populated
	case components.ChatContentItemsTypeFile:
		// chatContentItems.ChatContentFile is populated
}
```
