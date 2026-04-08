# ChatRequestIgnore


## Supported Types

### ProviderName

```go
chatRequestIgnore := components.CreateChatRequestIgnoreProviderName(components.ProviderName{/* values here */})
```

### 

```go
chatRequestIgnore := components.CreateChatRequestIgnoreStr(string{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch chatRequestIgnore.Type {
	case components.ChatRequestIgnoreTypeProviderName:
		// chatRequestIgnore.ProviderName is populated
	case components.ChatRequestIgnoreTypeStr:
		// chatRequestIgnore.Str is populated
}
```
