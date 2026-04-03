# InputsUnion

Input for a response request - can be a string or array of items


## Supported Types

### 

```go
inputsUnion := components.CreateInputsUnionStr(string{/* values here */})
```

### 

```go
inputsUnion := components.CreateInputsUnionArrayOfInputsUnion1([]components.InputsUnion1{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch inputsUnion.Type {
	case components.InputsUnionTypeStr:
		// inputsUnion.Str is populated
	case components.InputsUnionTypeArrayOfInputsUnion1:
		// inputsUnion.ArrayOfInputsUnion1 is populated
}
```
