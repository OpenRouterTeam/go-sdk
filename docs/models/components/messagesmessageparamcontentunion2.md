# MessagesMessageParamContentUnion2


## Supported Types

### 

```go
messagesMessageParamContentUnion2 := components.CreateMessagesMessageParamContentUnion2Str(string{/* values here */})
```

### 

```go
messagesMessageParamContentUnion2 := components.CreateMessagesMessageParamContentUnion2ArrayOfMessagesMessageParamContentUnion1([]components.MessagesMessageParamContentUnion1{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch messagesMessageParamContentUnion2.Type {
	case components.MessagesMessageParamContentUnion2TypeStr:
		// messagesMessageParamContentUnion2.Str is populated
	case components.MessagesMessageParamContentUnion2TypeArrayOfMessagesMessageParamContentUnion1:
		// messagesMessageParamContentUnion2.ArrayOfMessagesMessageParamContentUnion1 is populated
}
```
