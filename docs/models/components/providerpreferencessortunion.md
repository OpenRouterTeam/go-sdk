# ProviderPreferencesSortUnion

The sorting strategy to use for this request, if "order" is not specified. When set, no load balancing is performed.


## Supported Types

### ProviderPreferencesProviderSort

```go
providerPreferencesSortUnion := components.CreateProviderPreferencesSortUnionProviderPreferencesProviderSort(components.ProviderPreferencesProviderSort{/* values here */})
```

### ProviderPreferencesProviderSortConfigUnion

```go
providerPreferencesSortUnion := components.CreateProviderPreferencesSortUnionProviderPreferencesProviderSortConfigUnion(components.ProviderPreferencesProviderSortConfigUnion{/* values here */})
```

### ProviderPreferencesSortEnum

```go
providerPreferencesSortUnion := components.CreateProviderPreferencesSortUnionProviderPreferencesSortEnum(components.ProviderPreferencesSortEnum{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch providerPreferencesSortUnion.Type {
	case components.ProviderPreferencesSortUnionTypeProviderPreferencesProviderSort:
		// providerPreferencesSortUnion.ProviderPreferencesProviderSort is populated
	case components.ProviderPreferencesSortUnionTypeProviderPreferencesProviderSortConfigUnion:
		// providerPreferencesSortUnion.ProviderPreferencesProviderSortConfigUnion is populated
	case components.ProviderPreferencesSortUnionTypeProviderPreferencesSortEnum:
		// providerPreferencesSortUnion.ProviderPreferencesSortEnum is populated
	default:
		// Unknown type - use providerPreferencesSortUnion.GetUnknownRaw() for raw JSON
}
```
