# InputsUnion1


## Supported Types

### ReasoningItem

```go
inputsUnion1 := components.CreateInputsUnion1ReasoningItem(components.ReasoningItem{/* values here */})
```

### EasyInputMessage

```go
inputsUnion1 := components.CreateInputsUnion1EasyInputMessage(components.EasyInputMessage{/* values here */})
```

### InputMessageItem

```go
inputsUnion1 := components.CreateInputsUnion1InputMessageItem(components.InputMessageItem{/* values here */})
```

### FunctionCallItem

```go
inputsUnion1 := components.CreateInputsUnion1FunctionCallItem(components.FunctionCallItem{/* values here */})
```

### FunctionCallOutputItem

```go
inputsUnion1 := components.CreateInputsUnion1FunctionCallOutputItem(components.FunctionCallOutputItem{/* values here */})
```

### InputsMessage

```go
inputsUnion1 := components.CreateInputsUnion1InputsMessage(components.InputsMessage{/* values here */})
```

### InputsReasoning

```go
inputsUnion1 := components.CreateInputsUnion1InputsReasoning(components.InputsReasoning{/* values here */})
```

### OutputFunctionCallItem

```go
inputsUnion1 := components.CreateInputsUnion1OutputFunctionCallItem(components.OutputFunctionCallItem{/* values here */})
```

### OutputWebSearchCallItem

```go
inputsUnion1 := components.CreateInputsUnion1OutputWebSearchCallItem(components.OutputWebSearchCallItem{/* values here */})
```

### OutputFileSearchCallItem

```go
inputsUnion1 := components.CreateInputsUnion1OutputFileSearchCallItem(components.OutputFileSearchCallItem{/* values here */})
```

### OutputImageGenerationCallItem

```go
inputsUnion1 := components.CreateInputsUnion1OutputImageGenerationCallItem(components.OutputImageGenerationCallItem{/* values here */})
```

### OutputCodeInterpreterCallItem

```go
inputsUnion1 := components.CreateInputsUnion1OutputCodeInterpreterCallItem(components.OutputCodeInterpreterCallItem{/* values here */})
```

### OutputComputerCallItem

```go
inputsUnion1 := components.CreateInputsUnion1OutputComputerCallItem(components.OutputComputerCallItem{/* values here */})
```

### OutputDatetimeItem

```go
inputsUnion1 := components.CreateInputsUnion1OutputDatetimeItem(components.OutputDatetimeItem{/* values here */})
```

### OutputWebSearchServerToolItem

```go
inputsUnion1 := components.CreateInputsUnion1OutputWebSearchServerToolItem(components.OutputWebSearchServerToolItem{/* values here */})
```

### OutputCodeInterpreterServerToolItem

```go
inputsUnion1 := components.CreateInputsUnion1OutputCodeInterpreterServerToolItem(components.OutputCodeInterpreterServerToolItem{/* values here */})
```

### OutputFileSearchServerToolItem

```go
inputsUnion1 := components.CreateInputsUnion1OutputFileSearchServerToolItem(components.OutputFileSearchServerToolItem{/* values here */})
```

### OutputImageGenerationServerToolItem

```go
inputsUnion1 := components.CreateInputsUnion1OutputImageGenerationServerToolItem(components.OutputImageGenerationServerToolItem{/* values here */})
```

### OutputBrowserUseServerToolItem

```go
inputsUnion1 := components.CreateInputsUnion1OutputBrowserUseServerToolItem(components.OutputBrowserUseServerToolItem{/* values here */})
```

### OutputBashServerToolItem

```go
inputsUnion1 := components.CreateInputsUnion1OutputBashServerToolItem(components.OutputBashServerToolItem{/* values here */})
```

### OutputTextEditorServerToolItem

```go
inputsUnion1 := components.CreateInputsUnion1OutputTextEditorServerToolItem(components.OutputTextEditorServerToolItem{/* values here */})
```

### OutputApplyPatchServerToolItem

```go
inputsUnion1 := components.CreateInputsUnion1OutputApplyPatchServerToolItem(components.OutputApplyPatchServerToolItem{/* values here */})
```

### OutputWebFetchServerToolItem

```go
inputsUnion1 := components.CreateInputsUnion1OutputWebFetchServerToolItem(components.OutputWebFetchServerToolItem{/* values here */})
```

### OutputToolSearchServerToolItem

```go
inputsUnion1 := components.CreateInputsUnion1OutputToolSearchServerToolItem(components.OutputToolSearchServerToolItem{/* values here */})
```

### OutputMemoryServerToolItem

```go
inputsUnion1 := components.CreateInputsUnion1OutputMemoryServerToolItem(components.OutputMemoryServerToolItem{/* values here */})
```

### OutputMcpServerToolItem

```go
inputsUnion1 := components.CreateInputsUnion1OutputMcpServerToolItem(components.OutputMcpServerToolItem{/* values here */})
```

### OutputSearchModelsServerToolItem

```go
inputsUnion1 := components.CreateInputsUnion1OutputSearchModelsServerToolItem(components.OutputSearchModelsServerToolItem{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch inputsUnion1.Type {
	case components.InputsUnion1TypeReasoningItem:
		// inputsUnion1.ReasoningItem is populated
	case components.InputsUnion1TypeEasyInputMessage:
		// inputsUnion1.EasyInputMessage is populated
	case components.InputsUnion1TypeInputMessageItem:
		// inputsUnion1.InputMessageItem is populated
	case components.InputsUnion1TypeFunctionCallItem:
		// inputsUnion1.FunctionCallItem is populated
	case components.InputsUnion1TypeFunctionCallOutputItem:
		// inputsUnion1.FunctionCallOutputItem is populated
	case components.InputsUnion1TypeInputsMessage:
		// inputsUnion1.InputsMessage is populated
	case components.InputsUnion1TypeInputsReasoning:
		// inputsUnion1.InputsReasoning is populated
	case components.InputsUnion1TypeOutputFunctionCallItem:
		// inputsUnion1.OutputFunctionCallItem is populated
	case components.InputsUnion1TypeOutputWebSearchCallItem:
		// inputsUnion1.OutputWebSearchCallItem is populated
	case components.InputsUnion1TypeOutputFileSearchCallItem:
		// inputsUnion1.OutputFileSearchCallItem is populated
	case components.InputsUnion1TypeOutputImageGenerationCallItem:
		// inputsUnion1.OutputImageGenerationCallItem is populated
	case components.InputsUnion1TypeOutputCodeInterpreterCallItem:
		// inputsUnion1.OutputCodeInterpreterCallItem is populated
	case components.InputsUnion1TypeOutputComputerCallItem:
		// inputsUnion1.OutputComputerCallItem is populated
	case components.InputsUnion1TypeOutputDatetimeItem:
		// inputsUnion1.OutputDatetimeItem is populated
	case components.InputsUnion1TypeOutputWebSearchServerToolItem:
		// inputsUnion1.OutputWebSearchServerToolItem is populated
	case components.InputsUnion1TypeOutputCodeInterpreterServerToolItem:
		// inputsUnion1.OutputCodeInterpreterServerToolItem is populated
	case components.InputsUnion1TypeOutputFileSearchServerToolItem:
		// inputsUnion1.OutputFileSearchServerToolItem is populated
	case components.InputsUnion1TypeOutputImageGenerationServerToolItem:
		// inputsUnion1.OutputImageGenerationServerToolItem is populated
	case components.InputsUnion1TypeOutputBrowserUseServerToolItem:
		// inputsUnion1.OutputBrowserUseServerToolItem is populated
	case components.InputsUnion1TypeOutputBashServerToolItem:
		// inputsUnion1.OutputBashServerToolItem is populated
	case components.InputsUnion1TypeOutputTextEditorServerToolItem:
		// inputsUnion1.OutputTextEditorServerToolItem is populated
	case components.InputsUnion1TypeOutputApplyPatchServerToolItem:
		// inputsUnion1.OutputApplyPatchServerToolItem is populated
	case components.InputsUnion1TypeOutputWebFetchServerToolItem:
		// inputsUnion1.OutputWebFetchServerToolItem is populated
	case components.InputsUnion1TypeOutputToolSearchServerToolItem:
		// inputsUnion1.OutputToolSearchServerToolItem is populated
	case components.InputsUnion1TypeOutputMemoryServerToolItem:
		// inputsUnion1.OutputMemoryServerToolItem is populated
	case components.InputsUnion1TypeOutputMcpServerToolItem:
		// inputsUnion1.OutputMcpServerToolItem is populated
	case components.InputsUnion1TypeOutputSearchModelsServerToolItem:
		// inputsUnion1.OutputSearchModelsServerToolItem is populated
}
```
