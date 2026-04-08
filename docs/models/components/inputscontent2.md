# InputsContent2


## Supported Types

### 

```go
inputsContent2 := components.CreateInputsContent2ArrayOfInputsContent1([]components.InputsContent1{/* values here */})
```

### 

```go
inputsContent2 := components.CreateInputsContent2Str(string{/* values here */})
```

### 

```go
inputsContent2 := components.CreateInputsContent2Any(any{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch inputsContent2.Type {
	case components.InputsContent2TypeArrayOfInputsContent1:
		// inputsContent2.ArrayOfInputsContent1 is populated
	case components.InputsContent2TypeStr:
		// inputsContent2.Str is populated
	case components.InputsContent2TypeAny:
		// inputsContent2.Any is populated
}
```
