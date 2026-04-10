# Format


## Supported Types

### FormatText

```go
format := components.CreateFormatText(components.FormatText{/* values here */})
```

### FormatGrammar

```go
format := components.CreateFormatGrammar(components.FormatGrammar{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch format.Type {
	case components.FormatTypeTextValue:
		// format.FormatText is populated
	case components.FormatTypeGrammarValue:
		// format.FormatGrammar is populated
	default:
		// Unknown type - use format.GetUnknownRaw() for raw JSON
}
```
