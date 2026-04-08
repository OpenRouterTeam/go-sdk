# CreateResponsesResponse


## Supported Types

### OpenResponsesResult

```go
createResponsesResponse := operations.CreateCreateResponsesResponseOpenResponsesResult(components.OpenResponsesResult{/* values here */})
```

### 

```go
createResponsesResponse := operations.CreateCreateResponsesResponseEventStream(*stream.EventStream[operations.CreateResponsesResponseBody]{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch createResponsesResponse.Type {
	case operations.CreateResponsesResponseTypeOpenResponsesResult:
		// createResponsesResponse.OpenResponsesResult is populated
	case operations.CreateResponsesResponseTypeEventStream:
		// createResponsesResponse.EventStream is populated
}
```
