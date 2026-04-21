# InputUnion

The input to the generation — either a prompt string or an array of messages


## Supported Types

### Input1

```go
inputUnion := components.CreateInputUnionInput1(components.Input1{/* values here */})
```

### Input2

```go
inputUnion := components.CreateInputUnionInput2(components.Input2{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch inputUnion.Type {
	case components.InputUnionTypeInput1:
		// inputUnion.Input1 is populated
	case components.InputUnionTypeInput2:
		// inputUnion.Input2 is populated
}
```
