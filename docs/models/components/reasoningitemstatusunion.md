# ReasoningItemStatusUnion


## Supported Types

### ReasoningItemStatusCompleted

```go
reasoningItemStatusUnion := components.CreateReasoningItemStatusUnionReasoningItemStatusCompleted(components.ReasoningItemStatusCompleted{/* values here */})
```

### ReasoningItemStatusIncomplete

```go
reasoningItemStatusUnion := components.CreateReasoningItemStatusUnionReasoningItemStatusIncomplete(components.ReasoningItemStatusIncomplete{/* values here */})
```

### ReasoningItemStatusInProgress

```go
reasoningItemStatusUnion := components.CreateReasoningItemStatusUnionReasoningItemStatusInProgress(components.ReasoningItemStatusInProgress{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch reasoningItemStatusUnion.Type {
	case components.ReasoningItemStatusUnionTypeReasoningItemStatusCompleted:
		// reasoningItemStatusUnion.ReasoningItemStatusCompleted is populated
	case components.ReasoningItemStatusUnionTypeReasoningItemStatusIncomplete:
		// reasoningItemStatusUnion.ReasoningItemStatusIncomplete is populated
	case components.ReasoningItemStatusUnionTypeReasoningItemStatusInProgress:
		// reasoningItemStatusUnion.ReasoningItemStatusInProgress is populated
}
```
