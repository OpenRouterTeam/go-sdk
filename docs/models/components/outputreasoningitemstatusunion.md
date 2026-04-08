# OutputReasoningItemStatusUnion


## Supported Types

### OutputReasoningItemStatusCompleted

```go
outputReasoningItemStatusUnion := components.CreateOutputReasoningItemStatusUnionOutputReasoningItemStatusCompleted(components.OutputReasoningItemStatusCompleted{/* values here */})
```

### OutputReasoningItemStatusIncomplete

```go
outputReasoningItemStatusUnion := components.CreateOutputReasoningItemStatusUnionOutputReasoningItemStatusIncomplete(components.OutputReasoningItemStatusIncomplete{/* values here */})
```

### OutputReasoningItemStatusInProgress

```go
outputReasoningItemStatusUnion := components.CreateOutputReasoningItemStatusUnionOutputReasoningItemStatusInProgress(components.OutputReasoningItemStatusInProgress{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch outputReasoningItemStatusUnion.Type {
	case components.OutputReasoningItemStatusUnionTypeOutputReasoningItemStatusCompleted:
		// outputReasoningItemStatusUnion.OutputReasoningItemStatusCompleted is populated
	case components.OutputReasoningItemStatusUnionTypeOutputReasoningItemStatusIncomplete:
		// outputReasoningItemStatusUnion.OutputReasoningItemStatusIncomplete is populated
	case components.OutputReasoningItemStatusUnionTypeOutputReasoningItemStatusInProgress:
		// outputReasoningItemStatusUnion.OutputReasoningItemStatusInProgress is populated
}
```
