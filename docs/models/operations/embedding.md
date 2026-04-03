# Embedding


## Supported Types

### 

```go
embedding := operations.CreateEmbeddingArrayOfNumber([]float64{/* values here */})
```

### 

```go
embedding := operations.CreateEmbeddingStr(string{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch embedding.Type {
	case operations.EmbeddingTypeArrayOfNumber:
		// embedding.ArrayOfNumber is populated
	case operations.EmbeddingTypeStr:
		// embedding.Str is populated
}
```
