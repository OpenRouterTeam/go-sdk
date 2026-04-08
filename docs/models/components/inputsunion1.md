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

### OutputDatetimeItem

```go
inputsUnion1 := components.CreateInputsUnion1OutputDatetimeItem(components.OutputDatetimeItem{/* values here */})
```

### OutputServerToolItem

```go
inputsUnion1 := components.CreateInputsUnion1OutputServerToolItem(components.OutputServerToolItem{/* values here */})
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
	case components.InputsUnion1TypeOutputDatetimeItem:
		// inputsUnion1.OutputDatetimeItem is populated
	case components.InputsUnion1TypeOutputServerToolItem:
		// inputsUnion1.OutputServerToolItem is populated
}
```
