# ResponseFormat

Response format configuration


## Supported Types

### ChatFormatGrammarConfig

```go
responseFormat := components.CreateResponseFormatGrammar(components.ChatFormatGrammarConfig{/* values here */})
```

### FormatJSONObjectConfig

```go
responseFormat := components.CreateResponseFormatJSONObject(components.FormatJSONObjectConfig{/* values here */})
```

### ChatFormatJSONSchemaConfig

```go
responseFormat := components.CreateResponseFormatJSONSchema(components.ChatFormatJSONSchemaConfig{/* values here */})
```

### ChatFormatPythonConfig

```go
responseFormat := components.CreateResponseFormatPython(components.ChatFormatPythonConfig{/* values here */})
```

### ChatFormatTextConfig

```go
responseFormat := components.CreateResponseFormatText(components.ChatFormatTextConfig{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch responseFormat.Type {
	case components.ResponseFormatTypeGrammar:
		// responseFormat.ChatFormatGrammarConfig is populated
	case components.ResponseFormatTypeJSONObject:
		// responseFormat.FormatJSONObjectConfig is populated
	case components.ResponseFormatTypeJSONSchema:
		// responseFormat.ChatFormatJSONSchemaConfig is populated
	case components.ResponseFormatTypePython:
		// responseFormat.ChatFormatPythonConfig is populated
	case components.ResponseFormatTypeText:
		// responseFormat.ChatFormatTextConfig is populated
}
```
