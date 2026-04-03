# InputUnion


## Supported Types

### 

```go
inputUnion := operations.CreateInputUnionStr(string{/* values here */})
```

### 

```go
inputUnion := operations.CreateInputUnionArrayOfStr([]string{/* values here */})
```

### 

```go
inputUnion := operations.CreateInputUnionArrayOfNumber([]float64{/* values here */})
```

### 

```go
inputUnion := operations.CreateInputUnionArrayOfArrayOfNumber([][]float64{/* values here */})
```

### 

```go
inputUnion := operations.CreateInputUnionArrayOfInput([]operations.Input{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch inputUnion.Type {
	case operations.InputUnionTypeStr:
		// inputUnion.Str is populated
	case operations.InputUnionTypeArrayOfStr:
		// inputUnion.ArrayOfStr is populated
	case operations.InputUnionTypeArrayOfNumber:
		// inputUnion.ArrayOfNumber is populated
	case operations.InputUnionTypeArrayOfArrayOfNumber:
		// inputUnion.ArrayOfArrayOfNumber is populated
	case operations.InputUnionTypeArrayOfInput:
		// inputUnion.ArrayOfInput is populated
}
```
