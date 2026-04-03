# ReasoningDetailUnion

Reasoning detail union schema


## Supported Types

### ReasoningDetailSummary

```go
reasoningDetailUnion := components.CreateReasoningDetailUnionReasoningSummary(components.ReasoningDetailSummary{/* values here */})
```

### ReasoningDetailEncrypted

```go
reasoningDetailUnion := components.CreateReasoningDetailUnionReasoningEncrypted(components.ReasoningDetailEncrypted{/* values here */})
```

### ReasoningDetailText

```go
reasoningDetailUnion := components.CreateReasoningDetailUnionReasoningText(components.ReasoningDetailText{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch reasoningDetailUnion.Type {
	case components.ReasoningDetailUnionTypeReasoningSummary:
		// reasoningDetailUnion.ReasoningDetailSummary is populated
	case components.ReasoningDetailUnionTypeReasoningEncrypted:
		// reasoningDetailUnion.ReasoningDetailEncrypted is populated
	case components.ReasoningDetailUnionTypeReasoningText:
		// reasoningDetailUnion.ReasoningDetailText is populated
}
```
