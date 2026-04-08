# FiltersUnion


## Supported Types

### Filters

```go
filtersUnion := components.CreateFiltersUnionFilters(components.Filters{/* values here */})
```

### CompoundFilter

```go
filtersUnion := components.CreateFiltersUnionCompoundFilter(components.CompoundFilter{/* values here */})
```

### 

```go
filtersUnion := components.CreateFiltersUnionAny(any{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch filtersUnion.Type {
	case components.FiltersUnionTypeFilters:
		// filtersUnion.Filters is populated
	case components.FiltersUnionTypeCompoundFilter:
		// filtersUnion.CompoundFilter is populated
	case components.FiltersUnionTypeAny:
		// filtersUnion.Any is populated
}
```
