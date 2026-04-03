# ChatRequestOrder


## Supported Types

### ProviderName

```go
chatRequestOrder := components.CreateChatRequestOrderProviderName(components.ProviderName{/* values here */})
```

### 

```go
chatRequestOrder := components.CreateChatRequestOrderStr(string{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch chatRequestOrder.Type {
	case components.ChatRequestOrderTypeProviderName:
		// chatRequestOrder.ProviderName is populated
	case components.ChatRequestOrderTypeStr:
		// chatRequestOrder.Str is populated
}
```
