# OutputItems

An output item from the response


## Supported Types

### OutputCodeInterpreterCallItem

```go
outputItems := components.CreateOutputItemsCodeInterpreterCall(components.OutputCodeInterpreterCallItem{/* values here */})
```

### OutputComputerCallItem

```go
outputItems := components.CreateOutputItemsComputerCall(components.OutputComputerCallItem{/* values here */})
```

### OutputFileSearchCallItem

```go
outputItems := components.CreateOutputItemsFileSearchCall(components.OutputFileSearchCallItem{/* values here */})
```

### OutputFunctionCallItem

```go
outputItems := components.CreateOutputItemsFunctionCall(components.OutputFunctionCallItem{/* values here */})
```

### OutputImageGenerationCallItem

```go
outputItems := components.CreateOutputItemsImageGenerationCall(components.OutputImageGenerationCallItem{/* values here */})
```

### OutputMessageItem

```go
outputItems := components.CreateOutputItemsMessage(components.OutputMessageItem{/* values here */})
```

### OutputApplyPatchServerToolItem

```go
outputItems := components.CreateOutputItemsOpenrouterApplyPatch(components.OutputApplyPatchServerToolItem{/* values here */})
```

### OutputBashServerToolItem

```go
outputItems := components.CreateOutputItemsOpenrouterBash(components.OutputBashServerToolItem{/* values here */})
```

### OutputBrowserUseServerToolItem

```go
outputItems := components.CreateOutputItemsOpenrouterBrowserUse(components.OutputBrowserUseServerToolItem{/* values here */})
```

### OutputCodeInterpreterServerToolItem

```go
outputItems := components.CreateOutputItemsOpenrouterCodeInterpreter(components.OutputCodeInterpreterServerToolItem{/* values here */})
```

### OutputDatetimeItem

```go
outputItems := components.CreateOutputItemsOpenrouterDatetime(components.OutputDatetimeItem{/* values here */})
```

### OutputFileSearchServerToolItem

```go
outputItems := components.CreateOutputItemsOpenrouterFileSearch(components.OutputFileSearchServerToolItem{/* values here */})
```

### OutputImageGenerationServerToolItem

```go
outputItems := components.CreateOutputItemsOpenrouterImageGeneration(components.OutputImageGenerationServerToolItem{/* values here */})
```

### OutputMcpServerToolItem

```go
outputItems := components.CreateOutputItemsOpenrouterMcp(components.OutputMcpServerToolItem{/* values here */})
```

### OutputMemoryServerToolItem

```go
outputItems := components.CreateOutputItemsOpenrouterMemory(components.OutputMemoryServerToolItem{/* values here */})
```

### OutputTextEditorServerToolItem

```go
outputItems := components.CreateOutputItemsOpenrouterTextEditor(components.OutputTextEditorServerToolItem{/* values here */})
```

### OutputToolSearchServerToolItem

```go
outputItems := components.CreateOutputItemsOpenrouterToolSearch(components.OutputToolSearchServerToolItem{/* values here */})
```

### OutputWebFetchServerToolItem

```go
outputItems := components.CreateOutputItemsOpenrouterWebFetch(components.OutputWebFetchServerToolItem{/* values here */})
```

### OutputWebSearchServerToolItem

```go
outputItems := components.CreateOutputItemsOpenrouterWebSearch(components.OutputWebSearchServerToolItem{/* values here */})
```

### OutputReasoningItem

```go
outputItems := components.CreateOutputItemsReasoning(components.OutputReasoningItem{/* values here */})
```

### OutputWebSearchCallItem

```go
outputItems := components.CreateOutputItemsWebSearchCall(components.OutputWebSearchCallItem{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch outputItems.Type {
	case components.OutputItemsTypeCodeInterpreterCall:
		// outputItems.OutputCodeInterpreterCallItem is populated
	case components.OutputItemsTypeComputerCall:
		// outputItems.OutputComputerCallItem is populated
	case components.OutputItemsTypeFileSearchCall:
		// outputItems.OutputFileSearchCallItem is populated
	case components.OutputItemsTypeFunctionCall:
		// outputItems.OutputFunctionCallItem is populated
	case components.OutputItemsTypeImageGenerationCall:
		// outputItems.OutputImageGenerationCallItem is populated
	case components.OutputItemsTypeMessage:
		// outputItems.OutputMessageItem is populated
	case components.OutputItemsTypeOpenrouterApplyPatch:
		// outputItems.OutputApplyPatchServerToolItem is populated
	case components.OutputItemsTypeOpenrouterBash:
		// outputItems.OutputBashServerToolItem is populated
	case components.OutputItemsTypeOpenrouterBrowserUse:
		// outputItems.OutputBrowserUseServerToolItem is populated
	case components.OutputItemsTypeOpenrouterCodeInterpreter:
		// outputItems.OutputCodeInterpreterServerToolItem is populated
	case components.OutputItemsTypeOpenrouterDatetime:
		// outputItems.OutputDatetimeItem is populated
	case components.OutputItemsTypeOpenrouterFileSearch:
		// outputItems.OutputFileSearchServerToolItem is populated
	case components.OutputItemsTypeOpenrouterImageGeneration:
		// outputItems.OutputImageGenerationServerToolItem is populated
	case components.OutputItemsTypeOpenrouterMcp:
		// outputItems.OutputMcpServerToolItem is populated
	case components.OutputItemsTypeOpenrouterMemory:
		// outputItems.OutputMemoryServerToolItem is populated
	case components.OutputItemsTypeOpenrouterTextEditor:
		// outputItems.OutputTextEditorServerToolItem is populated
	case components.OutputItemsTypeOpenrouterToolSearch:
		// outputItems.OutputToolSearchServerToolItem is populated
	case components.OutputItemsTypeOpenrouterWebFetch:
		// outputItems.OutputWebFetchServerToolItem is populated
	case components.OutputItemsTypeOpenrouterWebSearch:
		// outputItems.OutputWebSearchServerToolItem is populated
	case components.OutputItemsTypeReasoning:
		// outputItems.OutputReasoningItem is populated
	case components.OutputItemsTypeWebSearchCall:
		// outputItems.OutputWebSearchCallItem is populated
	default:
		// Unknown type - use outputItems.GetUnknownRaw() for raw JSON
}
```
