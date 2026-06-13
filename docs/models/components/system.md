# System


## Supported Types

### 

```go
system := components.CreateSystemStr(string{/* values here */})
```

### 

```go
system := components.CreateSystemArrayOfAnthropicTextBlockParam([]components.AnthropicTextBlockParam{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch system.Type {
	case components.SystemTypeStr:
		// system.Str is populated
	case components.SystemTypeArrayOfAnthropicTextBlockParam:
		// system.ArrayOfAnthropicTextBlockParam is populated
}
```
