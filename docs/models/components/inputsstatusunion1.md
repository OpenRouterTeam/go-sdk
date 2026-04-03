# InputsStatusUnion1


## Supported Types

### InputsStatusCompleted1

```go
inputsStatusUnion1 := components.CreateInputsStatusUnion1InputsStatusCompleted1(components.InputsStatusCompleted1{/* values here */})
```

### InputsStatusIncomplete1

```go
inputsStatusUnion1 := components.CreateInputsStatusUnion1InputsStatusIncomplete1(components.InputsStatusIncomplete1{/* values here */})
```

### InputsStatusInProgress1

```go
inputsStatusUnion1 := components.CreateInputsStatusUnion1InputsStatusInProgress1(components.InputsStatusInProgress1{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch inputsStatusUnion1.Type {
	case components.InputsStatusUnion1TypeInputsStatusCompleted1:
		// inputsStatusUnion1.InputsStatusCompleted1 is populated
	case components.InputsStatusUnion1TypeInputsStatusIncomplete1:
		// inputsStatusUnion1.InputsStatusIncomplete1 is populated
	case components.InputsStatusUnion1TypeInputsStatusInProgress1:
		// inputsStatusUnion1.InputsStatusInProgress1 is populated
}
```
