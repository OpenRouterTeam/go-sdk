# BaseInputsRoleUnion2


## Supported Types

### BaseInputsRoleUser2

```go
baseInputsRoleUnion2 := components.CreateBaseInputsRoleUnion2BaseInputsRoleUser2(components.BaseInputsRoleUser2{/* values here */})
```

### BaseInputsRoleSystem2

```go
baseInputsRoleUnion2 := components.CreateBaseInputsRoleUnion2BaseInputsRoleSystem2(components.BaseInputsRoleSystem2{/* values here */})
```

### BaseInputsRoleDeveloper2

```go
baseInputsRoleUnion2 := components.CreateBaseInputsRoleUnion2BaseInputsRoleDeveloper2(components.BaseInputsRoleDeveloper2{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch baseInputsRoleUnion2.Type {
	case components.BaseInputsRoleUnion2TypeBaseInputsRoleUser2:
		// baseInputsRoleUnion2.BaseInputsRoleUser2 is populated
	case components.BaseInputsRoleUnion2TypeBaseInputsRoleSystem2:
		// baseInputsRoleUnion2.BaseInputsRoleSystem2 is populated
	case components.BaseInputsRoleUnion2TypeBaseInputsRoleDeveloper2:
		// baseInputsRoleUnion2.BaseInputsRoleDeveloper2 is populated
}
```
