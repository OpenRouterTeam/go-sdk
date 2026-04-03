# ChatRequestProviderSortConfigUnion

The provider sorting strategy (price, throughput, latency)


## Supported Types

### ChatRequestProviderSortConfig

```go
chatRequestProviderSortConfigUnion := components.CreateChatRequestProviderSortConfigUnionChatRequestProviderSortConfig(components.ChatRequestProviderSortConfig{/* values here */})
```

### ChatRequestProviderSortConfigEnum

```go
chatRequestProviderSortConfigUnion := components.CreateChatRequestProviderSortConfigUnionChatRequestProviderSortConfigEnum(components.ChatRequestProviderSortConfigEnum{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch chatRequestProviderSortConfigUnion.Type {
	case components.ChatRequestProviderSortConfigUnionTypeChatRequestProviderSortConfig:
		// chatRequestProviderSortConfigUnion.ChatRequestProviderSortConfig is populated
	case components.ChatRequestProviderSortConfigUnionTypeChatRequestProviderSortConfigEnum:
		// chatRequestProviderSortConfigUnion.ChatRequestProviderSortConfigEnum is populated
}
```
