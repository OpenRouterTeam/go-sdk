# BaseInputsContent2


## Supported Types

### 

```go
baseInputsContent2 := components.CreateBaseInputsContent2ArrayOfBaseInputsContent1([]components.BaseInputsContent1{/* values here */})
```

### 

```go
baseInputsContent2 := components.CreateBaseInputsContent2Str(string{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch baseInputsContent2.Type {
	case components.BaseInputsContent2TypeArrayOfBaseInputsContent1:
		// baseInputsContent2.ArrayOfBaseInputsContent1 is populated
	case components.BaseInputsContent2TypeStr:
		// baseInputsContent2.Str is populated
}
```
