# Ignore


## Supported Types

### ProviderName

```go
ignore := components.CreateIgnoreProviderName(components.ProviderName{/* values here */})
```

### 

```go
ignore := components.CreateIgnoreStr(string{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch ignore.Type {
	case components.IgnoreTypeProviderName:
		// ignore.ProviderName is populated
	case components.IgnoreTypeStr:
		// ignore.Str is populated
}
```
