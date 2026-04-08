# ChatRequestOnly


## Supported Types

### ProviderName

```go
chatRequestOnly := components.CreateChatRequestOnlyProviderName(components.ProviderName{/* values here */})
```

### 

```go
chatRequestOnly := components.CreateChatRequestOnlyStr(string{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch chatRequestOnly.Type {
	case components.ChatRequestOnlyTypeProviderName:
		// chatRequestOnly.ProviderName is populated
	case components.ChatRequestOnlyTypeStr:
		// chatRequestOnly.Str is populated
}
```
