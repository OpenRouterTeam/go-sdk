# ResponsesRequestOnly


## Supported Types

### ProviderName

```go
responsesRequestOnly := components.CreateResponsesRequestOnlyProviderName(components.ProviderName{/* values here */})
```

### 

```go
responsesRequestOnly := components.CreateResponsesRequestOnlyStr(string{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch responsesRequestOnly.Type {
	case components.ResponsesRequestOnlyTypeProviderName:
		// responsesRequestOnly.ProviderName is populated
	case components.ResponsesRequestOnlyTypeStr:
		// responsesRequestOnly.Str is populated
}
```
