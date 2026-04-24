# ChatFunctionTool

Tool definition for function calling (regular function or OpenRouter built-in server tool)


## Supported Types

### ChatFunctionToolFunction

```go
chatFunctionTool := components.CreateChatFunctionToolChatFunctionToolFunction(components.ChatFunctionToolFunction{/* values here */})
```

### DatetimeServerTool

```go
chatFunctionTool := components.CreateChatFunctionToolDatetimeServerTool(components.DatetimeServerTool{/* values here */})
```

### ImageGenerationServerToolOpenRouter

```go
chatFunctionTool := components.CreateChatFunctionToolImageGenerationServerToolOpenRouter(components.ImageGenerationServerToolOpenRouter{/* values here */})
```

### ChatSearchModelsServerTool

```go
chatFunctionTool := components.CreateChatFunctionToolChatSearchModelsServerTool(components.ChatSearchModelsServerTool{/* values here */})
```

### WebFetchServerTool

```go
chatFunctionTool := components.CreateChatFunctionToolWebFetchServerTool(components.WebFetchServerTool{/* values here */})
```

### OpenRouterWebSearchServerTool

```go
chatFunctionTool := components.CreateChatFunctionToolOpenRouterWebSearchServerTool(components.OpenRouterWebSearchServerTool{/* values here */})
```

### ChatWebSearchShorthand

```go
chatFunctionTool := components.CreateChatFunctionToolChatWebSearchShorthand(components.ChatWebSearchShorthand{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch chatFunctionTool.Type {
	case components.ChatFunctionToolUnionTypeChatFunctionToolFunction:
		// chatFunctionTool.ChatFunctionToolFunction is populated
	case components.ChatFunctionToolUnionTypeDatetimeServerTool:
		// chatFunctionTool.DatetimeServerTool is populated
	case components.ChatFunctionToolUnionTypeImageGenerationServerToolOpenRouter:
		// chatFunctionTool.ImageGenerationServerToolOpenRouter is populated
	case components.ChatFunctionToolUnionTypeChatSearchModelsServerTool:
		// chatFunctionTool.ChatSearchModelsServerTool is populated
	case components.ChatFunctionToolUnionTypeWebFetchServerTool:
		// chatFunctionTool.WebFetchServerTool is populated
	case components.ChatFunctionToolUnionTypeOpenRouterWebSearchServerTool:
		// chatFunctionTool.OpenRouterWebSearchServerTool is populated
	case components.ChatFunctionToolUnionTypeChatWebSearchShorthand:
		// chatFunctionTool.ChatWebSearchShorthand is populated
}
```
