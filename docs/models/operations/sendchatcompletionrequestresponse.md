# SendChatCompletionRequestResponse


## Supported Types

### ChatResult

```go
sendChatCompletionRequestResponse := operations.CreateSendChatCompletionRequestResponseChatResult(components.ChatResult{/* values here */})
```

### 

```go
sendChatCompletionRequestResponse := operations.CreateSendChatCompletionRequestResponseEventStream(*stream.EventStream[operations.SendChatCompletionRequestResponseBody]{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch sendChatCompletionRequestResponse.Type {
	case operations.SendChatCompletionRequestResponseTypeChatResult:
		// sendChatCompletionRequestResponse.ChatResult is populated
	case operations.SendChatCompletionRequestResponseTypeEventStream:
		// sendChatCompletionRequestResponse.EventStream is populated
}
```
