# CreateRerankResponse


## Supported Types

### CreateRerankResponseBody

```go
createRerankResponse := operations.CreateCreateRerankResponseCreateRerankResponseBody(operations.CreateRerankResponseBody{/* values here */})
```

### 

```go
createRerankResponse := operations.CreateCreateRerankResponseStr(string{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch createRerankResponse.Type {
	case operations.CreateRerankResponseTypeCreateRerankResponseBody:
		// createRerankResponse.CreateRerankResponseBody is populated
	case operations.CreateRerankResponseTypeStr:
		// createRerankResponse.Str is populated
}
```
