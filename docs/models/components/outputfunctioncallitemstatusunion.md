# OutputFunctionCallItemStatusUnion


## Supported Types

### OutputFunctionCallItemStatusCompleted

```go
outputFunctionCallItemStatusUnion := components.CreateOutputFunctionCallItemStatusUnionOutputFunctionCallItemStatusCompleted(components.OutputFunctionCallItemStatusCompleted{/* values here */})
```

### OutputFunctionCallItemStatusIncomplete

```go
outputFunctionCallItemStatusUnion := components.CreateOutputFunctionCallItemStatusUnionOutputFunctionCallItemStatusIncomplete(components.OutputFunctionCallItemStatusIncomplete{/* values here */})
```

### OutputFunctionCallItemStatusInProgress

```go
outputFunctionCallItemStatusUnion := components.CreateOutputFunctionCallItemStatusUnionOutputFunctionCallItemStatusInProgress(components.OutputFunctionCallItemStatusInProgress{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch outputFunctionCallItemStatusUnion.Type {
	case components.OutputFunctionCallItemStatusUnionTypeOutputFunctionCallItemStatusCompleted:
		// outputFunctionCallItemStatusUnion.OutputFunctionCallItemStatusCompleted is populated
	case components.OutputFunctionCallItemStatusUnionTypeOutputFunctionCallItemStatusIncomplete:
		// outputFunctionCallItemStatusUnion.OutputFunctionCallItemStatusIncomplete is populated
	case components.OutputFunctionCallItemStatusUnionTypeOutputFunctionCallItemStatusInProgress:
		// outputFunctionCallItemStatusUnion.OutputFunctionCallItemStatusInProgress is populated
}
```
