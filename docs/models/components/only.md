# Only


## Supported Types

### ProviderName

```go
only := components.CreateOnlyProviderName(components.ProviderName{/* values here */})
```

### 

```go
only := components.CreateOnlyStr(string{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch only.Type {
	case components.OnlyTypeProviderName:
		// only.ProviderName is populated
	case components.OnlyTypeStr:
		// only.Str is populated
}
```
