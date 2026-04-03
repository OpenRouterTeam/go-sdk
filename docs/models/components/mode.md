# Mode


## Supported Types

### ModeAuto

```go
mode := components.CreateModeModeAuto(components.ModeAuto{/* values here */})
```

### ModeRequired

```go
mode := components.CreateModeModeRequired(components.ModeRequired{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch mode.Type {
	case components.ModeTypeModeAuto:
		// mode.ModeAuto is populated
	case components.ModeTypeModeRequired:
		// mode.ModeRequired is populated
}
```
