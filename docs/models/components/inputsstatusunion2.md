# InputsStatusUnion2


## Supported Types

### InputsStatusCompleted2

```go
inputsStatusUnion2 := components.CreateInputsStatusUnion2InputsStatusCompleted2(components.InputsStatusCompleted2{/* values here */})
```

### InputsStatusIncomplete2

```go
inputsStatusUnion2 := components.CreateInputsStatusUnion2InputsStatusIncomplete2(components.InputsStatusIncomplete2{/* values here */})
```

### InputsStatusInProgress2

```go
inputsStatusUnion2 := components.CreateInputsStatusUnion2InputsStatusInProgress2(components.InputsStatusInProgress2{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch inputsStatusUnion2.Type {
	case components.InputsStatusUnion2TypeInputsStatusCompleted2:
		// inputsStatusUnion2.InputsStatusCompleted2 is populated
	case components.InputsStatusUnion2TypeInputsStatusIncomplete2:
		// inputsStatusUnion2.InputsStatusIncomplete2 is populated
	case components.InputsStatusUnion2TypeInputsStatusInProgress2:
		// inputsStatusUnion2.InputsStatusInProgress2 is populated
}
```
