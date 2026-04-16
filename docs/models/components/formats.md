# Formats

Text response format configuration


## Supported Types

### FormatTextConfig

```go
formats := components.CreateFormatsText(components.FormatTextConfig{/* values here */})
```

### FormatJSONObjectConfig

```go
formats := components.CreateFormatsJSONObject(components.FormatJSONObjectConfig{/* values here */})
```

### FormatJSONSchemaConfig

```go
formats := components.CreateFormatsJSONSchema(components.FormatJSONSchemaConfig{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch formats.Type {
	case components.FormatsTypeText:
		// formats.FormatTextConfig is populated
	case components.FormatsTypeJSONObject:
		// formats.FormatJSONObjectConfig is populated
	case components.FormatsTypeJSONSchema:
		// formats.FormatJSONSchemaConfig is populated
	default:
		// Unknown type - use formats.GetUnknownRaw() for raw JSON
}
```
