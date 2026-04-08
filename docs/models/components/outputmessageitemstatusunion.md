# OutputMessageItemStatusUnion


## Supported Types

### OutputMessageItemStatusCompleted

```go
outputMessageItemStatusUnion := components.CreateOutputMessageItemStatusUnionOutputMessageItemStatusCompleted(components.OutputMessageItemStatusCompleted{/* values here */})
```

### OutputMessageItemStatusIncomplete

```go
outputMessageItemStatusUnion := components.CreateOutputMessageItemStatusUnionOutputMessageItemStatusIncomplete(components.OutputMessageItemStatusIncomplete{/* values here */})
```

### OutputMessageItemStatusInProgress

```go
outputMessageItemStatusUnion := components.CreateOutputMessageItemStatusUnionOutputMessageItemStatusInProgress(components.OutputMessageItemStatusInProgress{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch outputMessageItemStatusUnion.Type {
	case components.OutputMessageItemStatusUnionTypeOutputMessageItemStatusCompleted:
		// outputMessageItemStatusUnion.OutputMessageItemStatusCompleted is populated
	case components.OutputMessageItemStatusUnionTypeOutputMessageItemStatusIncomplete:
		// outputMessageItemStatusUnion.OutputMessageItemStatusIncomplete is populated
	case components.OutputMessageItemStatusUnionTypeOutputMessageItemStatusInProgress:
		// outputMessageItemStatusUnion.OutputMessageItemStatusInProgress is populated
}
```
