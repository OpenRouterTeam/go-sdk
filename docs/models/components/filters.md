# Filters


## Supported Types

### FileSearchServerToolFilters

```go
filters := components.CreateFiltersFileSearchServerToolFilters(components.FileSearchServerToolFilters{/* values here */})
```

### CompoundFilter

```go
filters := components.CreateFiltersCompoundFilter(components.CompoundFilter{/* values here */})
```

### 

```go
filters := components.CreateFiltersAny(any{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch filters.Type {
	case components.FiltersUnionTypeFileSearchServerToolFilters:
		// filters.FileSearchServerToolFilters is populated
	case components.FiltersUnionTypeCompoundFilter:
		// filters.CompoundFilter is populated
	case components.FiltersUnionTypeAny:
		// filters.Any is populated
}
```
