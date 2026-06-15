# Value1

Filter value (scalar or array depending on operator)


## Supported Types

### 

```go
value1 := operations.CreateValue1Str(string{/* values here */})
```

### 

```go
value1 := operations.CreateValue1Number(float64{/* values here */})
```

### 

```go
value1 := operations.CreateValue1ArrayOfValue2([]operations.Value2{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch value1.Type {
	case operations.Value1TypeStr:
		// value1.Str is populated
	case operations.Value1TypeNumber:
		// value1.Number is populated
	case operations.Value1TypeArrayOfValue2:
		// value1.ArrayOfValue2 is populated
}
```
