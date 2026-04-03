# ChatDeveloperMessageContent

Developer message content


## Supported Types

### 

```go
chatDeveloperMessageContent := components.CreateChatDeveloperMessageContentStr(string{/* values here */})
```

### 

```go
chatDeveloperMessageContent := components.CreateChatDeveloperMessageContentArrayOfChatContentText([]components.ChatContentText{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch chatDeveloperMessageContent.Type {
	case components.ChatDeveloperMessageContentTypeStr:
		// chatDeveloperMessageContent.Str is populated
	case components.ChatDeveloperMessageContentTypeArrayOfChatContentText:
		// chatDeveloperMessageContent.ArrayOfChatContentText is populated
}
```
