# Data


## Supported Types

### UnifiedBenchmarksAAItem

```go
data := components.CreateDataArtificialAnalysis(components.UnifiedBenchmarksAAItem{/* values here */})
```

### UnifiedBenchmarksDAItem

```go
data := components.CreateDataDesignArena(components.UnifiedBenchmarksDAItem{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch data.Type {
	case components.DataTypeArtificialAnalysis:
		// data.UnifiedBenchmarksAAItem is populated
	case components.DataTypeDesignArena:
		// data.UnifiedBenchmarksDAItem is populated
	default:
		// Unknown type - use data.GetUnknownRaw() for raw JSON
}
```
