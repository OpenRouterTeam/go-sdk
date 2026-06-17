# ToolChoice


## Supported Types

### ToolChoiceAuto

```go
toolChoice := components.CreateToolChoiceAuto(components.ToolChoiceAuto{/* values here */})
```

### ToolChoiceAny

```go
toolChoice := components.CreateToolChoiceAny(components.ToolChoiceAny{/* values here */})
```

### ToolChoiceNone

```go
toolChoice := components.CreateToolChoiceNone(components.ToolChoiceNone{/* values here */})
```

### ToolChoiceTool

```go
toolChoice := components.CreateToolChoiceTool(components.ToolChoiceTool{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch toolChoice.Type {
	case components.ToolChoiceTypeAutoValue:
		// toolChoice.ToolChoiceAuto is populated
	case components.ToolChoiceTypeAny:
		// toolChoice.ToolChoiceAny is populated
	case components.ToolChoiceTypeNone:
		// toolChoice.ToolChoiceNone is populated
	case components.ToolChoiceTypeTool:
		// toolChoice.ToolChoiceTool is populated
}
```
