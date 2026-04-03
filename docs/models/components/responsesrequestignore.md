# ResponsesRequestIgnore


## Supported Types

### ProviderName

```go
responsesRequestIgnore := components.CreateResponsesRequestIgnoreProviderName(components.ProviderName{/* values here */})
```

### 

```go
responsesRequestIgnore := components.CreateResponsesRequestIgnoreStr(string{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch responsesRequestIgnore.Type {
	case components.ResponsesRequestIgnoreTypeProviderName:
		// responsesRequestIgnore.ProviderName is populated
	case components.ResponsesRequestIgnoreTypeStr:
		// responsesRequestIgnore.Str is populated
}
```
