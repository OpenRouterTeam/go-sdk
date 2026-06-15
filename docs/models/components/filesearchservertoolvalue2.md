# FileSearchServerToolValue2


## Supported Types

### 

```go
fileSearchServerToolValue2 := components.CreateFileSearchServerToolValue2Str(string{/* values here */})
```

### 

```go
fileSearchServerToolValue2 := components.CreateFileSearchServerToolValue2Number(float64{/* values here */})
```

### 

```go
fileSearchServerToolValue2 := components.CreateFileSearchServerToolValue2Boolean(bool{/* values here */})
```

### 

```go
fileSearchServerToolValue2 := components.CreateFileSearchServerToolValue2ArrayOfFileSearchServerToolValue1([]components.FileSearchServerToolValue1{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch fileSearchServerToolValue2.Type {
	case components.FileSearchServerToolValue2TypeStr:
		// fileSearchServerToolValue2.Str is populated
	case components.FileSearchServerToolValue2TypeNumber:
		// fileSearchServerToolValue2.Number is populated
	case components.FileSearchServerToolValue2TypeBoolean:
		// fileSearchServerToolValue2.Boolean is populated
	case components.FileSearchServerToolValue2TypeArrayOfFileSearchServerToolValue1:
		// fileSearchServerToolValue2.ArrayOfFileSearchServerToolValue1 is populated
}
```
