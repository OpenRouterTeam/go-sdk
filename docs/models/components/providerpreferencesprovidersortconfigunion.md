# ProviderPreferencesProviderSortConfigUnion

The provider sorting strategy (price, throughput, latency)


## Supported Types

### ProviderPreferencesProviderSortConfig

```go
providerPreferencesProviderSortConfigUnion := components.CreateProviderPreferencesProviderSortConfigUnionProviderPreferencesProviderSortConfig(components.ProviderPreferencesProviderSortConfig{/* values here */})
```

### ProviderPreferencesProviderSortConfigEnum

```go
providerPreferencesProviderSortConfigUnion := components.CreateProviderPreferencesProviderSortConfigUnionProviderPreferencesProviderSortConfigEnum(components.ProviderPreferencesProviderSortConfigEnum{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch providerPreferencesProviderSortConfigUnion.Type {
	case components.ProviderPreferencesProviderSortConfigUnionTypeProviderPreferencesProviderSortConfig:
		// providerPreferencesProviderSortConfigUnion.ProviderPreferencesProviderSortConfig is populated
	case components.ProviderPreferencesProviderSortConfigUnionTypeProviderPreferencesProviderSortConfigEnum:
		// providerPreferencesProviderSortConfigUnion.ProviderPreferencesProviderSortConfigEnum is populated
}
```
