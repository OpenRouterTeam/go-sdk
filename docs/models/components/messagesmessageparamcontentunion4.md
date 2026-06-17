# MessagesMessageParamContentUnion4


## Supported Types

### AnthropicTextBlockParam

```go
messagesMessageParamContentUnion4 := components.CreateMessagesMessageParamContentUnion4Text(components.AnthropicTextBlockParam{/* values here */})
```

### AnthropicImageBlockParam

```go
messagesMessageParamContentUnion4 := components.CreateMessagesMessageParamContentUnion4Image(components.AnthropicImageBlockParam{/* values here */})
```

### AnthropicDocumentBlockParam

```go
messagesMessageParamContentUnion4 := components.CreateMessagesMessageParamContentUnion4Document(components.AnthropicDocumentBlockParam{/* values here */})
```

### ContentToolUse

```go
messagesMessageParamContentUnion4 := components.CreateMessagesMessageParamContentUnion4ToolUse(components.ContentToolUse{/* values here */})
```

### ContentToolResult

```go
messagesMessageParamContentUnion4 := components.CreateMessagesMessageParamContentUnion4ToolResult(components.ContentToolResult{/* values here */})
```

### ContentThinking

```go
messagesMessageParamContentUnion4 := components.CreateMessagesMessageParamContentUnion4Thinking(components.ContentThinking{/* values here */})
```

### ContentRedactedThinking

```go
messagesMessageParamContentUnion4 := components.CreateMessagesMessageParamContentUnion4RedactedThinking(components.ContentRedactedThinking{/* values here */})
```

### ContentServerToolUse

```go
messagesMessageParamContentUnion4 := components.CreateMessagesMessageParamContentUnion4ServerToolUse(components.ContentServerToolUse{/* values here */})
```

### ContentWebSearchToolResult

```go
messagesMessageParamContentUnion4 := components.CreateMessagesMessageParamContentUnion4WebSearchToolResult(components.ContentWebSearchToolResult{/* values here */})
```

### AnthropicSearchResultBlockParam

```go
messagesMessageParamContentUnion4 := components.CreateMessagesMessageParamContentUnion4SearchResult(components.AnthropicSearchResultBlockParam{/* values here */})
```

### ContentCompaction

```go
messagesMessageParamContentUnion4 := components.CreateMessagesMessageParamContentUnion4Compaction(components.ContentCompaction{/* values here */})
```

### MessagesAdvisorToolResultBlock

```go
messagesMessageParamContentUnion4 := components.CreateMessagesMessageParamContentUnion4AdvisorToolResult(components.MessagesAdvisorToolResultBlock{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch messagesMessageParamContentUnion4.Type {
	case components.MessagesMessageParamContentUnion4TypeText:
		// messagesMessageParamContentUnion4.AnthropicTextBlockParam is populated
	case components.MessagesMessageParamContentUnion4TypeImage:
		// messagesMessageParamContentUnion4.AnthropicImageBlockParam is populated
	case components.MessagesMessageParamContentUnion4TypeDocument:
		// messagesMessageParamContentUnion4.AnthropicDocumentBlockParam is populated
	case components.MessagesMessageParamContentUnion4TypeToolUse:
		// messagesMessageParamContentUnion4.ContentToolUse is populated
	case components.MessagesMessageParamContentUnion4TypeToolResult:
		// messagesMessageParamContentUnion4.ContentToolResult is populated
	case components.MessagesMessageParamContentUnion4TypeThinking:
		// messagesMessageParamContentUnion4.ContentThinking is populated
	case components.MessagesMessageParamContentUnion4TypeRedactedThinking:
		// messagesMessageParamContentUnion4.ContentRedactedThinking is populated
	case components.MessagesMessageParamContentUnion4TypeServerToolUse:
		// messagesMessageParamContentUnion4.ContentServerToolUse is populated
	case components.MessagesMessageParamContentUnion4TypeWebSearchToolResult:
		// messagesMessageParamContentUnion4.ContentWebSearchToolResult is populated
	case components.MessagesMessageParamContentUnion4TypeSearchResult:
		// messagesMessageParamContentUnion4.AnthropicSearchResultBlockParam is populated
	case components.MessagesMessageParamContentUnion4TypeCompaction:
		// messagesMessageParamContentUnion4.ContentCompaction is populated
	case components.MessagesMessageParamContentUnion4TypeAdvisorToolResult:
		// messagesMessageParamContentUnion4.MessagesAdvisorToolResultBlock is populated
}
```
