# BaseInputsRoleUnion


## Supported Types

### BaseInputsRoleUser

```go
baseInputsRoleUnion := components.CreateBaseInputsRoleUnionBaseInputsRoleUser(components.BaseInputsRoleUser{/* values here */})
```

### BaseInputsRoleSystem

```go
baseInputsRoleUnion := components.CreateBaseInputsRoleUnionBaseInputsRoleSystem(components.BaseInputsRoleSystem{/* values here */})
```

### BaseInputsRoleAssistant

```go
baseInputsRoleUnion := components.CreateBaseInputsRoleUnionBaseInputsRoleAssistant(components.BaseInputsRoleAssistant{/* values here */})
```

### BaseInputsRoleDeveloper

```go
baseInputsRoleUnion := components.CreateBaseInputsRoleUnionBaseInputsRoleDeveloper(components.BaseInputsRoleDeveloper{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch baseInputsRoleUnion.Type {
	case components.BaseInputsRoleUnionTypeBaseInputsRoleUser:
		// baseInputsRoleUnion.BaseInputsRoleUser is populated
	case components.BaseInputsRoleUnionTypeBaseInputsRoleSystem:
		// baseInputsRoleUnion.BaseInputsRoleSystem is populated
	case components.BaseInputsRoleUnionTypeBaseInputsRoleAssistant:
		// baseInputsRoleUnion.BaseInputsRoleAssistant is populated
	case components.BaseInputsRoleUnionTypeBaseInputsRoleDeveloper:
		// baseInputsRoleUnion.BaseInputsRoleDeveloper is populated
}
```
