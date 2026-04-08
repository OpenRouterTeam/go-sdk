# Order


## Supported Types

### ProviderName

```go
order := components.CreateOrderProviderName(components.ProviderName{/* values here */})
```

### 

```go
order := components.CreateOrderStr(string{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch order.Type {
	case components.OrderTypeProviderName:
		// order.ProviderName is populated
	case components.OrderTypeStr:
		// order.Str is populated
}
```
