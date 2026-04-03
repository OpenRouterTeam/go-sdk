# ProviderPreferencesIgnore


## Supported Types

### ProviderName

```go
providerPreferencesIgnore := components.CreateProviderPreferencesIgnoreProviderName(components.ProviderName{/* values here */})
```

### 

```go
providerPreferencesIgnore := components.CreateProviderPreferencesIgnoreStr(string{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch providerPreferencesIgnore.Type {
	case components.ProviderPreferencesIgnoreTypeProviderName:
		// providerPreferencesIgnore.ProviderName is populated
	case components.ProviderPreferencesIgnoreTypeStr:
		// providerPreferencesIgnore.Str is populated
}
```
