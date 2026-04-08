# Sort

The sorting strategy to use for this request, if "order" is not specified. When set, no load balancing is performed.


## Supported Types

### ProviderSort

```go
sort := components.CreateSortProviderSort(components.ProviderSort{/* values here */})
```

### ProviderSortConfig

```go
sort := components.CreateSortProviderSortConfig(components.ProviderSortConfig{/* values here */})
```

### 

```go
sort := components.CreateSortAny(any{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch sort.Type {
	case components.SortTypeProviderSort:
		// sort.ProviderSort is populated
	case components.SortTypeProviderSortConfig:
		// sort.ProviderSortConfig is populated
	case components.SortTypeAny:
		// sort.Any is populated
}
```
