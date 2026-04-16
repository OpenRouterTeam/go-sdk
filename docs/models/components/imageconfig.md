# ImageConfig


## Supported Types

### 

```go
imageConfig := components.CreateImageConfigStr(string{/* values here */})
```

### 

```go
imageConfig := components.CreateImageConfigNumber(float64{/* values here */})
```

### 

```go
imageConfig := components.CreateImageConfigArrayOfAny([]any{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch imageConfig.Type {
	case components.ImageConfigTypeStr:
		// imageConfig.Str is populated
	case components.ImageConfigTypeNumber:
		// imageConfig.Number is populated
	case components.ImageConfigTypeArrayOfAny:
		// imageConfig.ArrayOfAny is populated
}
```
