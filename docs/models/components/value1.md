# Value1


## Supported Types

### 

```go
value1 := components.CreateValue1Str(string{/* values here */})
```

### 

```go
value1 := components.CreateValue1Number(float64{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch value1.Type {
	case components.Value1TypeStr:
		// value1.Str is populated
	case components.Value1TypeNumber:
		// value1.Number is populated
}
```
