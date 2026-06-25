# UnifiedBenchmarksResponseData


## Supported Types

### UnifiedBenchmarksAAItem

```go
unifiedBenchmarksResponseData := components.CreateUnifiedBenchmarksResponseDataArtificialAnalysis(components.UnifiedBenchmarksAAItem{/* values here */})
```

### UnifiedBenchmarksDAItem

```go
unifiedBenchmarksResponseData := components.CreateUnifiedBenchmarksResponseDataDesignArena(components.UnifiedBenchmarksDAItem{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch unifiedBenchmarksResponseData.Type {
	case components.UnifiedBenchmarksResponseDataTypeArtificialAnalysis:
		// unifiedBenchmarksResponseData.UnifiedBenchmarksAAItem is populated
	case components.UnifiedBenchmarksResponseDataTypeDesignArena:
		// unifiedBenchmarksResponseData.UnifiedBenchmarksDAItem is populated
	default:
		// Unknown type - use unifiedBenchmarksResponseData.GetUnknownRaw() for raw JSON
}
```
