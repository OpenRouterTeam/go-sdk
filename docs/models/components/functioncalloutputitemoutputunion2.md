# FunctionCallOutputItemOutputUnion2


## Supported Types

### 

```go
functionCallOutputItemOutputUnion2 := components.CreateFunctionCallOutputItemOutputUnion2Str(string{/* values here */})
```

### 

```go
functionCallOutputItemOutputUnion2 := components.CreateFunctionCallOutputItemOutputUnion2ArrayOfFunctionCallOutputItemOutputUnion1([]components.FunctionCallOutputItemOutputUnion1{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch functionCallOutputItemOutputUnion2.Type {
	case components.FunctionCallOutputItemOutputUnion2TypeStr:
		// functionCallOutputItemOutputUnion2.Str is populated
	case components.FunctionCallOutputItemOutputUnion2TypeArrayOfFunctionCallOutputItemOutputUnion1:
		// functionCallOutputItemOutputUnion2.ArrayOfFunctionCallOutputItemOutputUnion1 is populated
}
```
