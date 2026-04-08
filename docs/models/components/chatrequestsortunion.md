# ChatRequestSortUnion

The sorting strategy to use for this request, if "order" is not specified. When set, no load balancing is performed.


## Supported Types

### ChatRequestProviderSort

```go
chatRequestSortUnion := components.CreateChatRequestSortUnionChatRequestProviderSort(components.ChatRequestProviderSort{/* values here */})
```

### ChatRequestProviderSortConfigUnion

```go
chatRequestSortUnion := components.CreateChatRequestSortUnionChatRequestProviderSortConfigUnion(components.ChatRequestProviderSortConfigUnion{/* values here */})
```

### ChatRequestSortEnum

```go
chatRequestSortUnion := components.CreateChatRequestSortUnionChatRequestSortEnum(components.ChatRequestSortEnum{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch chatRequestSortUnion.Type {
	case components.ChatRequestSortUnionTypeChatRequestProviderSort:
		// chatRequestSortUnion.ChatRequestProviderSort is populated
	case components.ChatRequestSortUnionTypeChatRequestProviderSortConfigUnion:
		// chatRequestSortUnion.ChatRequestProviderSortConfigUnion is populated
	case components.ChatRequestSortUnionTypeChatRequestSortEnum:
		// chatRequestSortUnion.ChatRequestSortEnum is populated
	default:
		// Unknown type - use chatRequestSortUnion.GetUnknownRaw() for raw JSON
}
```
