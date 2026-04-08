# OutputMessageStatusUnion


## Supported Types

### OutputMessageStatusCompleted

```go
outputMessageStatusUnion := components.CreateOutputMessageStatusUnionOutputMessageStatusCompleted(components.OutputMessageStatusCompleted{/* values here */})
```

### OutputMessageStatusIncomplete

```go
outputMessageStatusUnion := components.CreateOutputMessageStatusUnionOutputMessageStatusIncomplete(components.OutputMessageStatusIncomplete{/* values here */})
```

### OutputMessageStatusInProgress

```go
outputMessageStatusUnion := components.CreateOutputMessageStatusUnionOutputMessageStatusInProgress(components.OutputMessageStatusInProgress{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch outputMessageStatusUnion.Type {
	case components.OutputMessageStatusUnionTypeOutputMessageStatusCompleted:
		// outputMessageStatusUnion.OutputMessageStatusCompleted is populated
	case components.OutputMessageStatusUnionTypeOutputMessageStatusIncomplete:
		// outputMessageStatusUnion.OutputMessageStatusIncomplete is populated
	case components.OutputMessageStatusUnionTypeOutputMessageStatusInProgress:
		// outputMessageStatusUnion.OutputMessageStatusInProgress is populated
}
```
