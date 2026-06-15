# CustomToolCallOutputItemOutputUnion2


## Supported Types

### 

```go
customToolCallOutputItemOutputUnion2 := components.CreateCustomToolCallOutputItemOutputUnion2Str(string{/* values here */})
```

### 

```go
customToolCallOutputItemOutputUnion2 := components.CreateCustomToolCallOutputItemOutputUnion2ArrayOfCustomToolCallOutputItemOutputUnion1([]components.CustomToolCallOutputItemOutputUnion1{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch customToolCallOutputItemOutputUnion2.Type {
	case components.CustomToolCallOutputItemOutputUnion2TypeStr:
		// customToolCallOutputItemOutputUnion2.Str is populated
	case components.CustomToolCallOutputItemOutputUnion2TypeArrayOfCustomToolCallOutputItemOutputUnion1:
		// customToolCallOutputItemOutputUnion2.ArrayOfCustomToolCallOutputItemOutputUnion1 is populated
}
```
