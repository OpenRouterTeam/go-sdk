# ResponsesRequestPluginUnion


## Supported Types

### ResponsesRequestPluginAutoRouter

```go
responsesRequestPluginUnion := components.CreateResponsesRequestPluginUnionAutoRouter(components.ResponsesRequestPluginAutoRouter{/* values here */})
```

### ResponsesRequestPluginModeration

```go
responsesRequestPluginUnion := components.CreateResponsesRequestPluginUnionModeration(components.ResponsesRequestPluginModeration{/* values here */})
```

### ResponsesRequestPluginWeb

```go
responsesRequestPluginUnion := components.CreateResponsesRequestPluginUnionWeb(components.ResponsesRequestPluginWeb{/* values here */})
```

### ResponsesRequestPluginFileParser

```go
responsesRequestPluginUnion := components.CreateResponsesRequestPluginUnionFileParser(components.ResponsesRequestPluginFileParser{/* values here */})
```

### ResponsesRequestPluginResponseHealing

```go
responsesRequestPluginUnion := components.CreateResponsesRequestPluginUnionResponseHealing(components.ResponsesRequestPluginResponseHealing{/* values here */})
```

### ResponsesRequestPluginContextCompression

```go
responsesRequestPluginUnion := components.CreateResponsesRequestPluginUnionContextCompression(components.ResponsesRequestPluginContextCompression{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch responsesRequestPluginUnion.Type {
	case components.ResponsesRequestPluginUnionTypeAutoRouter:
		// responsesRequestPluginUnion.ResponsesRequestPluginAutoRouter is populated
	case components.ResponsesRequestPluginUnionTypeModeration:
		// responsesRequestPluginUnion.ResponsesRequestPluginModeration is populated
	case components.ResponsesRequestPluginUnionTypeWeb:
		// responsesRequestPluginUnion.ResponsesRequestPluginWeb is populated
	case components.ResponsesRequestPluginUnionTypeFileParser:
		// responsesRequestPluginUnion.ResponsesRequestPluginFileParser is populated
	case components.ResponsesRequestPluginUnionTypeResponseHealing:
		// responsesRequestPluginUnion.ResponsesRequestPluginResponseHealing is populated
	case components.ResponsesRequestPluginUnionTypeContextCompression:
		// responsesRequestPluginUnion.ResponsesRequestPluginContextCompression is populated
}
```
