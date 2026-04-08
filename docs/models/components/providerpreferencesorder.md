# ProviderPreferencesOrder


## Supported Types

### ProviderName

```go
providerPreferencesOrder := components.CreateProviderPreferencesOrderProviderName(components.ProviderName{/* values here */})
```

### 

```go
providerPreferencesOrder := components.CreateProviderPreferencesOrderStr(string{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch providerPreferencesOrder.Type {
	case components.ProviderPreferencesOrderTypeProviderName:
		// providerPreferencesOrder.ProviderName is populated
	case components.ProviderPreferencesOrderTypeStr:
		// providerPreferencesOrder.Str is populated
}
```
