# MessagesMessageParamContentUnion1


## Supported Types

### AnthropicTextBlockParam

```go
messagesMessageParamContentUnion1 := components.CreateMessagesMessageParamContentUnion1Text(components.AnthropicTextBlockParam{/* values here */})
```

### AnthropicImageBlockParam

```go
messagesMessageParamContentUnion1 := components.CreateMessagesMessageParamContentUnion1Image(components.AnthropicImageBlockParam{/* values here */})
```

### ContentToolReference

```go
messagesMessageParamContentUnion1 := components.CreateMessagesMessageParamContentUnion1ToolReference(components.ContentToolReference{/* values here */})
```

### AnthropicSearchResultBlockParam

```go
messagesMessageParamContentUnion1 := components.CreateMessagesMessageParamContentUnion1SearchResult(components.AnthropicSearchResultBlockParam{/* values here */})
```

### AnthropicDocumentBlockParam

```go
messagesMessageParamContentUnion1 := components.CreateMessagesMessageParamContentUnion1Document(components.AnthropicDocumentBlockParam{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch messagesMessageParamContentUnion1.Type {
	case components.MessagesMessageParamContentUnion1TypeText:
		// messagesMessageParamContentUnion1.AnthropicTextBlockParam is populated
	case components.MessagesMessageParamContentUnion1TypeImage:
		// messagesMessageParamContentUnion1.AnthropicImageBlockParam is populated
	case components.MessagesMessageParamContentUnion1TypeToolReference:
		// messagesMessageParamContentUnion1.ContentToolReference is populated
	case components.MessagesMessageParamContentUnion1TypeSearchResult:
		// messagesMessageParamContentUnion1.AnthropicSearchResultBlockParam is populated
	case components.MessagesMessageParamContentUnion1TypeDocument:
		// messagesMessageParamContentUnion1.AnthropicDocumentBlockParam is populated
}
```
