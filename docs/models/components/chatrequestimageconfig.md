# ChatRequestImageConfig


## Supported Types

### 

```go
chatRequestImageConfig := components.CreateChatRequestImageConfigStr(string{/* values here */})
```

### 

```go
chatRequestImageConfig := components.CreateChatRequestImageConfigNumber(float64{/* values here */})
```

### 

```go
chatRequestImageConfig := components.CreateChatRequestImageConfigArrayOfAny([]any{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch chatRequestImageConfig.Type {
	case components.ChatRequestImageConfigTypeStr:
		// chatRequestImageConfig.Str is populated
	case components.ChatRequestImageConfigTypeNumber:
		// chatRequestImageConfig.Number is populated
	case components.ChatRequestImageConfigTypeArrayOfAny:
		// chatRequestImageConfig.ArrayOfAny is populated
}
```
