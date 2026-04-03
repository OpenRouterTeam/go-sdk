# BaseInputsOutput2


## Supported Types

### 

```go
baseInputsOutput2 := components.CreateBaseInputsOutput2Str(string{/* values here */})
```

### 

```go
baseInputsOutput2 := components.CreateBaseInputsOutput2ArrayOfBaseInputsOutput1([]components.BaseInputsOutput1{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch baseInputsOutput2.Type {
	case components.BaseInputsOutput2TypeStr:
		// baseInputsOutput2.Str is populated
	case components.BaseInputsOutput2TypeArrayOfBaseInputsOutput1:
		// baseInputsOutput2.ArrayOfBaseInputsOutput1 is populated
}
```
