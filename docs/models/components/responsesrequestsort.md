# ResponsesRequestSort

The sorting strategy to use for this request, if "order" is not specified. When set, no load balancing is performed.


## Supported Types

### ProviderSort

```go
responsesRequestSort := components.CreateResponsesRequestSortProviderSort(components.ProviderSort{/* values here */})
```

### ProviderSortConfig

```go
responsesRequestSort := components.CreateResponsesRequestSortProviderSortConfig(components.ProviderSortConfig{/* values here */})
```

### 

```go
responsesRequestSort := components.CreateResponsesRequestSortAny(any{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch responsesRequestSort.Type {
	case components.ResponsesRequestSortTypeProviderSort:
		// responsesRequestSort.ProviderSort is populated
	case components.ResponsesRequestSortTypeProviderSortConfig:
		// responsesRequestSort.ProviderSortConfig is populated
	case components.ResponsesRequestSortTypeAny:
		// responsesRequestSort.Any is populated
}
```
