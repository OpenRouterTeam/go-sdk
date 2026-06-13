# Edit


## Supported Types

### EditClearToolUses20250919

```go
edit := components.CreateEditClearToolUses20250919(components.EditClearToolUses20250919{/* values here */})
```

### EditClearThinking20251015

```go
edit := components.CreateEditClearThinking20251015(components.EditClearThinking20251015{/* values here */})
```

### EditCompact20260112

```go
edit := components.CreateEditCompact20260112(components.EditCompact20260112{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch edit.Type {
	case components.EditTypeClearToolUses20250919:
		// edit.EditClearToolUses20250919 is populated
	case components.EditTypeClearThinking20251015:
		// edit.EditClearThinking20251015 is populated
	case components.EditTypeCompact20260112:
		// edit.EditCompact20260112 is populated
}
```
