# ProviderPreferencesOnly


## Supported Types

### ProviderName

```go
providerPreferencesOnly := components.CreateProviderPreferencesOnlyProviderName(components.ProviderName{/* values here */})
```

### 

```go
providerPreferencesOnly := components.CreateProviderPreferencesOnlyStr(string{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch providerPreferencesOnly.Type {
	case components.ProviderPreferencesOnlyTypeProviderName:
		// providerPreferencesOnly.ProviderName is populated
	case components.ProviderPreferencesOnlyTypeStr:
		// providerPreferencesOnly.Str is populated
}
```
