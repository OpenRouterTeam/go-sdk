# ChatContentItems1


## Supported Types

### LegacyChatContentVideo

```go
chatContentItems1 := components.CreateChatContentItems1InputVideo(components.LegacyChatContentVideo{/* values here */})
```

### ChatContentVideo

```go
chatContentItems1 := components.CreateChatContentItems1VideoURL(components.ChatContentVideo{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch chatContentItems1.Type {
	case components.ChatContentItems1TypeInputVideo:
		// chatContentItems1.LegacyChatContentVideo is populated
	case components.ChatContentItems1TypeVideoURL:
		// chatContentItems1.ChatContentVideo is populated
}
```
