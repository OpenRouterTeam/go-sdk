# BaseInputsUnion


## Supported Types

### 

```go
baseInputsUnion := components.CreateBaseInputsUnionStr(string{/* values here */})
```

### 

```go
baseInputsUnion := components.CreateBaseInputsUnionArrayOfBaseInputsUnion1([]components.BaseInputsUnion1{/* values here */})
```

### 

```go
baseInputsUnion := components.CreateBaseInputsUnionAny(any{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch baseInputsUnion.Type {
	case components.BaseInputsUnionTypeStr:
		// baseInputsUnion.Str is populated
	case components.BaseInputsUnionTypeArrayOfBaseInputsUnion1:
		// baseInputsUnion.ArrayOfBaseInputsUnion1 is populated
	case components.BaseInputsUnionTypeAny:
		// baseInputsUnion.Any is populated
}
```
