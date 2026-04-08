# AllowedToolsUnion


## Supported Types

### 

```go
allowedToolsUnion := components.CreateAllowedToolsUnionArrayOfStr([]string{/* values here */})
```

### AllowedTools

```go
allowedToolsUnion := components.CreateAllowedToolsUnionAllowedTools(components.AllowedTools{/* values here */})
```

### 

```go
allowedToolsUnion := components.CreateAllowedToolsUnionAny(any{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch allowedToolsUnion.Type {
	case components.AllowedToolsUnionTypeArrayOfStr:
		// allowedToolsUnion.ArrayOfStr is populated
	case components.AllowedToolsUnionTypeAllowedTools:
		// allowedToolsUnion.AllowedTools is populated
	case components.AllowedToolsUnionTypeAny:
		// allowedToolsUnion.Any is populated
}
```
