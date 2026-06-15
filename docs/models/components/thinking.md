# Thinking


## Supported Types

### ThinkingEnabled

```go
thinking := components.CreateThinkingEnabled(components.ThinkingEnabled{/* values here */})
```

### ThinkingDisabled

```go
thinking := components.CreateThinkingDisabled(components.ThinkingDisabled{/* values here */})
```

### ThinkingAdaptive

```go
thinking := components.CreateThinkingAdaptive(components.ThinkingAdaptive{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch thinking.Type {
	case components.ThinkingTypeEnabled:
		// thinking.ThinkingEnabled is populated
	case components.ThinkingTypeDisabled:
		// thinking.ThinkingDisabled is populated
	case components.ThinkingTypeAdaptive:
		// thinking.ThinkingAdaptive is populated
}
```
