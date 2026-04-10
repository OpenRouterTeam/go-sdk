# OpenResponsesResultToolUnion


## Supported Types

### OpenResponsesResultToolFunction

```go
openResponsesResultToolUnion := components.CreateOpenResponsesResultToolUnionFunction(components.OpenResponsesResultToolFunction{/* values here */})
```

### PreviewWebSearchServerTool

```go
openResponsesResultToolUnion := components.CreateOpenResponsesResultToolUnionWebSearchPreview(components.PreviewWebSearchServerTool{/* values here */})
```

### Preview20250311WebSearchServerTool

```go
openResponsesResultToolUnion := components.CreateOpenResponsesResultToolUnionWebSearchPreview20250311(components.Preview20250311WebSearchServerTool{/* values here */})
```

### LegacyWebSearchServerTool

```go
openResponsesResultToolUnion := components.CreateOpenResponsesResultToolUnionWebSearch(components.LegacyWebSearchServerTool{/* values here */})
```

### WebSearchServerTool

```go
openResponsesResultToolUnion := components.CreateOpenResponsesResultToolUnionWebSearch20250826(components.WebSearchServerTool{/* values here */})
```

### FileSearchServerTool

```go
openResponsesResultToolUnion := components.CreateOpenResponsesResultToolUnionFileSearch(components.FileSearchServerTool{/* values here */})
```

### ComputerUseServerTool

```go
openResponsesResultToolUnion := components.CreateOpenResponsesResultToolUnionComputerUsePreview(components.ComputerUseServerTool{/* values here */})
```

### CodeInterpreterServerTool

```go
openResponsesResultToolUnion := components.CreateOpenResponsesResultToolUnionCodeInterpreter(components.CodeInterpreterServerTool{/* values here */})
```

### McpServerTool

```go
openResponsesResultToolUnion := components.CreateOpenResponsesResultToolUnionMcp(components.McpServerTool{/* values here */})
```

### ImageGenerationServerTool

```go
openResponsesResultToolUnion := components.CreateOpenResponsesResultToolUnionImageGeneration(components.ImageGenerationServerTool{/* values here */})
```

### CodexLocalShellTool

```go
openResponsesResultToolUnion := components.CreateOpenResponsesResultToolUnionLocalShell(components.CodexLocalShellTool{/* values here */})
```

### ShellServerTool

```go
openResponsesResultToolUnion := components.CreateOpenResponsesResultToolUnionShell(components.ShellServerTool{/* values here */})
```

### ApplyPatchServerTool

```go
openResponsesResultToolUnion := components.CreateOpenResponsesResultToolUnionApplyPatch(components.ApplyPatchServerTool{/* values here */})
```

### CustomTool

```go
openResponsesResultToolUnion := components.CreateOpenResponsesResultToolUnionCustom(components.CustomTool{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch openResponsesResultToolUnion.Type {
	case components.OpenResponsesResultToolUnionTypeFunction:
		// openResponsesResultToolUnion.OpenResponsesResultToolFunction is populated
	case components.OpenResponsesResultToolUnionTypeWebSearchPreview:
		// openResponsesResultToolUnion.PreviewWebSearchServerTool is populated
	case components.OpenResponsesResultToolUnionTypeWebSearchPreview20250311:
		// openResponsesResultToolUnion.Preview20250311WebSearchServerTool is populated
	case components.OpenResponsesResultToolUnionTypeWebSearch:
		// openResponsesResultToolUnion.LegacyWebSearchServerTool is populated
	case components.OpenResponsesResultToolUnionTypeWebSearch20250826:
		// openResponsesResultToolUnion.WebSearchServerTool is populated
	case components.OpenResponsesResultToolUnionTypeFileSearch:
		// openResponsesResultToolUnion.FileSearchServerTool is populated
	case components.OpenResponsesResultToolUnionTypeComputerUsePreview:
		// openResponsesResultToolUnion.ComputerUseServerTool is populated
	case components.OpenResponsesResultToolUnionTypeCodeInterpreter:
		// openResponsesResultToolUnion.CodeInterpreterServerTool is populated
	case components.OpenResponsesResultToolUnionTypeMcp:
		// openResponsesResultToolUnion.McpServerTool is populated
	case components.OpenResponsesResultToolUnionTypeImageGeneration:
		// openResponsesResultToolUnion.ImageGenerationServerTool is populated
	case components.OpenResponsesResultToolUnionTypeLocalShell:
		// openResponsesResultToolUnion.CodexLocalShellTool is populated
	case components.OpenResponsesResultToolUnionTypeShell:
		// openResponsesResultToolUnion.ShellServerTool is populated
	case components.OpenResponsesResultToolUnionTypeApplyPatch:
		// openResponsesResultToolUnion.ApplyPatchServerTool is populated
	case components.OpenResponsesResultToolUnionTypeCustom:
		// openResponsesResultToolUnion.CustomTool is populated
	default:
		// Unknown type - use openResponsesResultToolUnion.GetUnknownRaw() for raw JSON
}
```
