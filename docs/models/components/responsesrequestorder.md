# ResponsesRequestOrder


## Supported Types

### ProviderName

```go
responsesRequestOrder := components.CreateResponsesRequestOrderProviderName(components.ProviderName{/* values here */})
```

### 

```go
responsesRequestOrder := components.CreateResponsesRequestOrderStr(string{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch responsesRequestOrder.Type {
	case components.ResponsesRequestOrderTypeProviderName:
		// responsesRequestOrder.ProviderName is populated
	case components.ResponsesRequestOrderTypeStr:
		// responsesRequestOrder.Str is populated
}
```
