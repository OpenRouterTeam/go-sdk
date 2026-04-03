# Value2


## Supported Types

### 

```go
value2 := components.CreateValue2Str(string{/* values here */})
```

### 

```go
value2 := components.CreateValue2Number(float64{/* values here */})
```

### 

```go
value2 := components.CreateValue2Boolean(bool{/* values here */})
```

### 

```go
value2 := components.CreateValue2ArrayOfValue1([]components.Value1{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch value2.Type {
	case components.Value2TypeStr:
		// value2.Str is populated
	case components.Value2TypeNumber:
		// value2.Number is populated
	case components.Value2TypeBoolean:
		// value2.Boolean is populated
	case components.Value2TypeArrayOfValue1:
		// value2.ArrayOfValue1 is populated
}
```
