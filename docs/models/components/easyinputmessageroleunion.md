# EasyInputMessageRoleUnion


## Supported Types

### EasyInputMessageRoleUser

```go
easyInputMessageRoleUnion := components.CreateEasyInputMessageRoleUnionEasyInputMessageRoleUser(components.EasyInputMessageRoleUser{/* values here */})
```

### EasyInputMessageRoleSystem

```go
easyInputMessageRoleUnion := components.CreateEasyInputMessageRoleUnionEasyInputMessageRoleSystem(components.EasyInputMessageRoleSystem{/* values here */})
```

### EasyInputMessageRoleAssistant

```go
easyInputMessageRoleUnion := components.CreateEasyInputMessageRoleUnionEasyInputMessageRoleAssistant(components.EasyInputMessageRoleAssistant{/* values here */})
```

### EasyInputMessageRoleDeveloper

```go
easyInputMessageRoleUnion := components.CreateEasyInputMessageRoleUnionEasyInputMessageRoleDeveloper(components.EasyInputMessageRoleDeveloper{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch easyInputMessageRoleUnion.Type {
	case components.EasyInputMessageRoleUnionTypeEasyInputMessageRoleUser:
		// easyInputMessageRoleUnion.EasyInputMessageRoleUser is populated
	case components.EasyInputMessageRoleUnionTypeEasyInputMessageRoleSystem:
		// easyInputMessageRoleUnion.EasyInputMessageRoleSystem is populated
	case components.EasyInputMessageRoleUnionTypeEasyInputMessageRoleAssistant:
		// easyInputMessageRoleUnion.EasyInputMessageRoleAssistant is populated
	case components.EasyInputMessageRoleUnionTypeEasyInputMessageRoleDeveloper:
		// easyInputMessageRoleUnion.EasyInputMessageRoleDeveloper is populated
}
```
