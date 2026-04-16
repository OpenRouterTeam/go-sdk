# ReasoningDetailUnion

Reasoning detail union schema


## Supported Types

### ReasoningDetailEncrypted

```go
reasoningDetailUnion := components.CreateReasoningDetailUnionReasoningEncrypted(components.ReasoningDetailEncrypted{/* values here */})
```

### ReasoningDetailSummary

```go
reasoningDetailUnion := components.CreateReasoningDetailUnionReasoningSummary(components.ReasoningDetailSummary{/* values here */})
```

### ReasoningDetailText

```go
reasoningDetailUnion := components.CreateReasoningDetailUnionReasoningText(components.ReasoningDetailText{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch reasoningDetailUnion.Type {
	case components.ReasoningDetailUnionTypeReasoningEncrypted:
		// reasoningDetailUnion.ReasoningDetailEncrypted is populated
	case components.ReasoningDetailUnionTypeReasoningSummary:
		// reasoningDetailUnion.ReasoningDetailSummary is populated
	case components.ReasoningDetailUnionTypeReasoningText:
		// reasoningDetailUnion.ReasoningDetailText is populated
	default:
		// Unknown type - use reasoningDetailUnion.GetUnknownRaw() for raw JSON
}
```
