# ChatSystemMessageContent

System message content


## Supported Types

### 

```go
chatSystemMessageContent := components.CreateChatSystemMessageContentStr(string{/* values here */})
```

### 

```go
chatSystemMessageContent := components.CreateChatSystemMessageContentArrayOfChatContentText([]components.ChatContentText{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch chatSystemMessageContent.Type {
	case components.ChatSystemMessageContentTypeStr:
		// chatSystemMessageContent.Str is populated
	case components.ChatSystemMessageContentTypeArrayOfChatContentText:
		// chatSystemMessageContent.ArrayOfChatContentText is populated
}
```
