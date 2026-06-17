# MessagesMessageParamContentUnion3


## Supported Types

### 

```go
messagesMessageParamContentUnion3 := components.CreateMessagesMessageParamContentUnion3ArrayOfAnthropicWebSearchResultBlockParam([]components.AnthropicWebSearchResultBlockParam{/* values here */})
```

### ContentWebSearchToolResultError

```go
messagesMessageParamContentUnion3 := components.CreateMessagesMessageParamContentUnion3ContentWebSearchToolResultError(components.ContentWebSearchToolResultError{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch messagesMessageParamContentUnion3.Type {
	case components.MessagesMessageParamContentUnion3TypeArrayOfAnthropicWebSearchResultBlockParam:
		// messagesMessageParamContentUnion3.ArrayOfAnthropicWebSearchResultBlockParam is populated
	case components.MessagesMessageParamContentUnion3TypeContentWebSearchToolResultError:
		// messagesMessageParamContentUnion3.ContentWebSearchToolResultError is populated
}
```
