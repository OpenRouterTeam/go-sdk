# InputMessageItemRoleUnion


## Supported Types

### InputMessageItemRoleUser

```go
inputMessageItemRoleUnion := components.CreateInputMessageItemRoleUnionInputMessageItemRoleUser(components.InputMessageItemRoleUser{/* values here */})
```

### InputMessageItemRoleSystem

```go
inputMessageItemRoleUnion := components.CreateInputMessageItemRoleUnionInputMessageItemRoleSystem(components.InputMessageItemRoleSystem{/* values here */})
```

### InputMessageItemRoleDeveloper

```go
inputMessageItemRoleUnion := components.CreateInputMessageItemRoleUnionInputMessageItemRoleDeveloper(components.InputMessageItemRoleDeveloper{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch inputMessageItemRoleUnion.Type {
	case components.InputMessageItemRoleUnionTypeInputMessageItemRoleUser:
		// inputMessageItemRoleUnion.InputMessageItemRoleUser is populated
	case components.InputMessageItemRoleUnionTypeInputMessageItemRoleSystem:
		// inputMessageItemRoleUnion.InputMessageItemRoleSystem is populated
	case components.InputMessageItemRoleUnionTypeInputMessageItemRoleDeveloper:
		// inputMessageItemRoleUnion.InputMessageItemRoleDeveloper is populated
}
```
