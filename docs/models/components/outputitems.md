# OutputItems

An output item from the response


## Supported Types

### OutputMessageItem

```go
outputItems := components.CreateOutputItemsOutputMessageItem(components.OutputMessageItem{/* values here */})
```

### OutputReasoningItem

```go
outputItems := components.CreateOutputItemsOutputReasoningItem(components.OutputReasoningItem{/* values here */})
```

### OutputFunctionCallItem

```go
outputItems := components.CreateOutputItemsOutputFunctionCallItem(components.OutputFunctionCallItem{/* values here */})
```

### OutputWebSearchCallItem

```go
outputItems := components.CreateOutputItemsOutputWebSearchCallItem(components.OutputWebSearchCallItem{/* values here */})
```

### OutputFileSearchCallItem

```go
outputItems := components.CreateOutputItemsOutputFileSearchCallItem(components.OutputFileSearchCallItem{/* values here */})
```

### OutputImageGenerationCallItem

```go
outputItems := components.CreateOutputItemsOutputImageGenerationCallItem(components.OutputImageGenerationCallItem{/* values here */})
```

### OutputServerToolItem

```go
outputItems := components.CreateOutputItemsOutputServerToolItem(components.OutputServerToolItem{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch outputItems.Type {
	case components.OutputItemsTypeOutputMessageItem:
		// outputItems.OutputMessageItem is populated
	case components.OutputItemsTypeOutputReasoningItem:
		// outputItems.OutputReasoningItem is populated
	case components.OutputItemsTypeOutputFunctionCallItem:
		// outputItems.OutputFunctionCallItem is populated
	case components.OutputItemsTypeOutputWebSearchCallItem:
		// outputItems.OutputWebSearchCallItem is populated
	case components.OutputItemsTypeOutputFileSearchCallItem:
		// outputItems.OutputFileSearchCallItem is populated
	case components.OutputItemsTypeOutputImageGenerationCallItem:
		// outputItems.OutputImageGenerationCallItem is populated
	case components.OutputItemsTypeOutputServerToolItem:
		// outputItems.OutputServerToolItem is populated
}
```
