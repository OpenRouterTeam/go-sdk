# ResponsesRequestToolUnion


## Supported Types

### ResponsesRequestToolFunction

```go
responsesRequestToolUnion := components.CreateResponsesRequestToolUnionFunction(components.ResponsesRequestToolFunction{/* values here */})
```

### PreviewWebSearchServerTool

```go
responsesRequestToolUnion := components.CreateResponsesRequestToolUnionWebSearchPreview(components.PreviewWebSearchServerTool{/* values here */})
```

### Preview20250311WebSearchServerTool

```go
responsesRequestToolUnion := components.CreateResponsesRequestToolUnionWebSearchPreview20250311(components.Preview20250311WebSearchServerTool{/* values here */})
```

### LegacyWebSearchServerTool

```go
responsesRequestToolUnion := components.CreateResponsesRequestToolUnionWebSearch(components.LegacyWebSearchServerTool{/* values here */})
```

### WebSearchServerTool

```go
responsesRequestToolUnion := components.CreateResponsesRequestToolUnionWebSearch20250826(components.WebSearchServerTool{/* values here */})
```

### FileSearchServerTool

```go
responsesRequestToolUnion := components.CreateResponsesRequestToolUnionFileSearch(components.FileSearchServerTool{/* values here */})
```

### ComputerUseServerTool

```go
responsesRequestToolUnion := components.CreateResponsesRequestToolUnionComputerUsePreview(components.ComputerUseServerTool{/* values here */})
```

### CodeInterpreterServerTool

```go
responsesRequestToolUnion := components.CreateResponsesRequestToolUnionCodeInterpreter(components.CodeInterpreterServerTool{/* values here */})
```

### McpServerTool

```go
responsesRequestToolUnion := components.CreateResponsesRequestToolUnionMcp(components.McpServerTool{/* values here */})
```

### ImageGenerationServerTool

```go
responsesRequestToolUnion := components.CreateResponsesRequestToolUnionImageGeneration(components.ImageGenerationServerTool{/* values here */})
```

### CodexLocalShellTool

```go
responsesRequestToolUnion := components.CreateResponsesRequestToolUnionLocalShell(components.CodexLocalShellTool{/* values here */})
```

### ShellServerTool

```go
responsesRequestToolUnion := components.CreateResponsesRequestToolUnionShell(components.ShellServerTool{/* values here */})
```

### ApplyPatchServerTool

```go
responsesRequestToolUnion := components.CreateResponsesRequestToolUnionApplyPatch(components.ApplyPatchServerTool{/* values here */})
```

### CustomTool

```go
responsesRequestToolUnion := components.CreateResponsesRequestToolUnionCustom(components.CustomTool{/* values here */})
```

### DatetimeServerTool

```go
responsesRequestToolUnion := components.CreateResponsesRequestToolUnionOpenrouterDatetime(components.DatetimeServerTool{/* values here */})
```

### ImageGenerationServerToolOpenRouter

```go
responsesRequestToolUnion := components.CreateResponsesRequestToolUnionOpenrouterImageGeneration(components.ImageGenerationServerToolOpenRouter{/* values here */})
```

### ChatSearchModelsServerTool

```go
responsesRequestToolUnion := components.CreateResponsesRequestToolUnionOpenrouterExperimentalSearchModels(components.ChatSearchModelsServerTool{/* values here */})
```

### WebSearchServerToolOpenRouter

```go
responsesRequestToolUnion := components.CreateResponsesRequestToolUnionOpenrouterWebSearch(components.WebSearchServerToolOpenRouter{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch responsesRequestToolUnion.Type {
	case components.ResponsesRequestToolUnionTypeFunction:
		// responsesRequestToolUnion.ResponsesRequestToolFunction is populated
	case components.ResponsesRequestToolUnionTypeWebSearchPreview:
		// responsesRequestToolUnion.PreviewWebSearchServerTool is populated
	case components.ResponsesRequestToolUnionTypeWebSearchPreview20250311:
		// responsesRequestToolUnion.Preview20250311WebSearchServerTool is populated
	case components.ResponsesRequestToolUnionTypeWebSearch:
		// responsesRequestToolUnion.LegacyWebSearchServerTool is populated
	case components.ResponsesRequestToolUnionTypeWebSearch20250826:
		// responsesRequestToolUnion.WebSearchServerTool is populated
	case components.ResponsesRequestToolUnionTypeFileSearch:
		// responsesRequestToolUnion.FileSearchServerTool is populated
	case components.ResponsesRequestToolUnionTypeComputerUsePreview:
		// responsesRequestToolUnion.ComputerUseServerTool is populated
	case components.ResponsesRequestToolUnionTypeCodeInterpreter:
		// responsesRequestToolUnion.CodeInterpreterServerTool is populated
	case components.ResponsesRequestToolUnionTypeMcp:
		// responsesRequestToolUnion.McpServerTool is populated
	case components.ResponsesRequestToolUnionTypeImageGeneration:
		// responsesRequestToolUnion.ImageGenerationServerTool is populated
	case components.ResponsesRequestToolUnionTypeLocalShell:
		// responsesRequestToolUnion.CodexLocalShellTool is populated
	case components.ResponsesRequestToolUnionTypeShell:
		// responsesRequestToolUnion.ShellServerTool is populated
	case components.ResponsesRequestToolUnionTypeApplyPatch:
		// responsesRequestToolUnion.ApplyPatchServerTool is populated
	case components.ResponsesRequestToolUnionTypeCustom:
		// responsesRequestToolUnion.CustomTool is populated
	case components.ResponsesRequestToolUnionTypeOpenrouterDatetime:
		// responsesRequestToolUnion.DatetimeServerTool is populated
	case components.ResponsesRequestToolUnionTypeOpenrouterImageGeneration:
		// responsesRequestToolUnion.ImageGenerationServerToolOpenRouter is populated
	case components.ResponsesRequestToolUnionTypeOpenrouterExperimentalSearchModels:
		// responsesRequestToolUnion.ChatSearchModelsServerTool is populated
	case components.ResponsesRequestToolUnionTypeOpenrouterWebSearch:
		// responsesRequestToolUnion.WebSearchServerToolOpenRouter is populated
}
```
