# ResponsesRequestPlugin


## Supported Types

### AutoRouterPlugin

```go
responsesRequestPlugin := components.CreateResponsesRequestPluginAutoRouter(components.AutoRouterPlugin{/* values here */})
```

### ModerationPlugin

```go
responsesRequestPlugin := components.CreateResponsesRequestPluginModeration(components.ModerationPlugin{/* values here */})
```

### WebSearchPlugin

```go
responsesRequestPlugin := components.CreateResponsesRequestPluginWeb(components.WebSearchPlugin{/* values here */})
```

### FileParserPlugin

```go
responsesRequestPlugin := components.CreateResponsesRequestPluginFileParser(components.FileParserPlugin{/* values here */})
```

### ResponseHealingPlugin

```go
responsesRequestPlugin := components.CreateResponsesRequestPluginResponseHealing(components.ResponseHealingPlugin{/* values here */})
```

### ContextCompressionPlugin

```go
responsesRequestPlugin := components.CreateResponsesRequestPluginContextCompression(components.ContextCompressionPlugin{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch responsesRequestPlugin.Type {
	case components.ResponsesRequestPluginTypeAutoRouter:
		// responsesRequestPlugin.AutoRouterPlugin is populated
	case components.ResponsesRequestPluginTypeModeration:
		// responsesRequestPlugin.ModerationPlugin is populated
	case components.ResponsesRequestPluginTypeWeb:
		// responsesRequestPlugin.WebSearchPlugin is populated
	case components.ResponsesRequestPluginTypeFileParser:
		// responsesRequestPlugin.FileParserPlugin is populated
	case components.ResponsesRequestPluginTypeResponseHealing:
		// responsesRequestPlugin.ResponseHealingPlugin is populated
	case components.ResponsesRequestPluginTypeContextCompression:
		// responsesRequestPlugin.ContextCompressionPlugin is populated
}
```
