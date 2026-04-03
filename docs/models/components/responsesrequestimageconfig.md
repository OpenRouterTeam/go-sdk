# ResponsesRequestImageConfig


## Supported Types

### 

```go
responsesRequestImageConfig := components.CreateResponsesRequestImageConfigStr(string{/* values here */})
```

### 

```go
responsesRequestImageConfig := components.CreateResponsesRequestImageConfigNumber(float64{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch responsesRequestImageConfig.Type {
	case components.ResponsesRequestImageConfigTypeStr:
		// responsesRequestImageConfig.Str is populated
	case components.ResponsesRequestImageConfigTypeNumber:
		// responsesRequestImageConfig.Number is populated
}
```
