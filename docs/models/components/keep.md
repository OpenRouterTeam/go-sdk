# Keep


## Supported Types

### AnthropicThinkingTurns

```go
keep := components.CreateKeepAnthropicThinkingTurns(components.AnthropicThinkingTurns{/* values here */})
```

### KeepAll

```go
keep := components.CreateKeepKeepAll(components.KeepAll{/* values here */})
```

### KeepEnum

```go
keep := components.CreateKeepKeepEnum(components.KeepEnum{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch keep.Type {
	case components.KeepUnionTypeAnthropicThinkingTurns:
		// keep.AnthropicThinkingTurns is populated
	case components.KeepUnionTypeKeepAll:
		// keep.KeepAll is populated
	case components.KeepUnionTypeKeepEnum:
		// keep.KeepEnum is populated
}
```
