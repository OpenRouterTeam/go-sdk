# ChatRequestPlugin


## Supported Types

### AutoRouterPlugin

```go
chatRequestPlugin := components.CreateChatRequestPluginAutoRouter(components.AutoRouterPlugin{/* values here */})
```

### ContextCompressionPlugin

```go
chatRequestPlugin := components.CreateChatRequestPluginContextCompression(components.ContextCompressionPlugin{/* values here */})
```

### FileParserPlugin

```go
chatRequestPlugin := components.CreateChatRequestPluginFileParser(components.FileParserPlugin{/* values here */})
```

### ModerationPlugin

```go
chatRequestPlugin := components.CreateChatRequestPluginModeration(components.ModerationPlugin{/* values here */})
```

### ParetoRouterPlugin

```go
chatRequestPlugin := components.CreateChatRequestPluginParetoRouter(components.ParetoRouterPlugin{/* values here */})
```

### ResponseHealingPlugin

```go
chatRequestPlugin := components.CreateChatRequestPluginResponseHealing(components.ResponseHealingPlugin{/* values here */})
```

### WebSearchPlugin

```go
chatRequestPlugin := components.CreateChatRequestPluginWeb(components.WebSearchPlugin{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch chatRequestPlugin.Type {
	case components.ChatRequestPluginTypeAutoRouter:
		// chatRequestPlugin.AutoRouterPlugin is populated
	case components.ChatRequestPluginTypeContextCompression:
		// chatRequestPlugin.ContextCompressionPlugin is populated
	case components.ChatRequestPluginTypeFileParser:
		// chatRequestPlugin.FileParserPlugin is populated
	case components.ChatRequestPluginTypeModeration:
		// chatRequestPlugin.ModerationPlugin is populated
	case components.ChatRequestPluginTypeParetoRouter:
		// chatRequestPlugin.ParetoRouterPlugin is populated
	case components.ChatRequestPluginTypeResponseHealing:
		// chatRequestPlugin.ResponseHealingPlugin is populated
	case components.ChatRequestPluginTypeWeb:
		// chatRequestPlugin.WebSearchPlugin is populated
}
```
