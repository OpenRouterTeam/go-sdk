# CreateEmbeddingsResponse


## Supported Types

### CreateEmbeddingsResponseBody

```go
createEmbeddingsResponse := operations.CreateCreateEmbeddingsResponseCreateEmbeddingsResponseBody(operations.CreateEmbeddingsResponseBody{/* values here */})
```

### 

```go
createEmbeddingsResponse := operations.CreateCreateEmbeddingsResponseStr(string{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch createEmbeddingsResponse.Type {
	case operations.CreateEmbeddingsResponseTypeCreateEmbeddingsResponseBody:
		// createEmbeddingsResponse.CreateEmbeddingsResponseBody is populated
	case operations.CreateEmbeddingsResponseTypeStr:
		// createEmbeddingsResponse.Str is populated
}
```
