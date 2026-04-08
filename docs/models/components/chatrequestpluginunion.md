# ChatRequestPluginUnion


## Supported Types

### ChatRequestPluginAutoRouter

```go
chatRequestPluginUnion := components.CreateChatRequestPluginUnionAutoRouter(components.ChatRequestPluginAutoRouter{/* values here */})
```

### ChatRequestPluginModeration

```go
chatRequestPluginUnion := components.CreateChatRequestPluginUnionModeration(components.ChatRequestPluginModeration{/* values here */})
```

### ChatRequestPluginWeb

```go
chatRequestPluginUnion := components.CreateChatRequestPluginUnionWeb(components.ChatRequestPluginWeb{/* values here */})
```

### ChatRequestPluginFileParser

```go
chatRequestPluginUnion := components.CreateChatRequestPluginUnionFileParser(components.ChatRequestPluginFileParser{/* values here */})
```

### ChatRequestPluginResponseHealing

```go
chatRequestPluginUnion := components.CreateChatRequestPluginUnionResponseHealing(components.ChatRequestPluginResponseHealing{/* values here */})
```

### ChatRequestPluginContextCompression

```go
chatRequestPluginUnion := components.CreateChatRequestPluginUnionContextCompression(components.ChatRequestPluginContextCompression{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch chatRequestPluginUnion.Type {
	case components.ChatRequestPluginUnionTypeAutoRouter:
		// chatRequestPluginUnion.ChatRequestPluginAutoRouter is populated
	case components.ChatRequestPluginUnionTypeModeration:
		// chatRequestPluginUnion.ChatRequestPluginModeration is populated
	case components.ChatRequestPluginUnionTypeWeb:
		// chatRequestPluginUnion.ChatRequestPluginWeb is populated
	case components.ChatRequestPluginUnionTypeFileParser:
		// chatRequestPluginUnion.ChatRequestPluginFileParser is populated
	case components.ChatRequestPluginUnionTypeResponseHealing:
		// chatRequestPluginUnion.ChatRequestPluginResponseHealing is populated
	case components.ChatRequestPluginUnionTypeContextCompression:
		// chatRequestPluginUnion.ChatRequestPluginContextCompression is populated
}
```
