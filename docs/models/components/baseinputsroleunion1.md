# BaseInputsRoleUnion1


## Supported Types

### BaseInputsRoleUser1

```go
baseInputsRoleUnion1 := components.CreateBaseInputsRoleUnion1BaseInputsRoleUser1(components.BaseInputsRoleUser1{/* values here */})
```

### BaseInputsRoleSystem1

```go
baseInputsRoleUnion1 := components.CreateBaseInputsRoleUnion1BaseInputsRoleSystem1(components.BaseInputsRoleSystem1{/* values here */})
```

### BaseInputsRoleAssistant

```go
baseInputsRoleUnion1 := components.CreateBaseInputsRoleUnion1BaseInputsRoleAssistant(components.BaseInputsRoleAssistant{/* values here */})
```

### BaseInputsRoleDeveloper1

```go
baseInputsRoleUnion1 := components.CreateBaseInputsRoleUnion1BaseInputsRoleDeveloper1(components.BaseInputsRoleDeveloper1{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch baseInputsRoleUnion1.Type {
	case components.BaseInputsRoleUnion1TypeBaseInputsRoleUser1:
		// baseInputsRoleUnion1.BaseInputsRoleUser1 is populated
	case components.BaseInputsRoleUnion1TypeBaseInputsRoleSystem1:
		// baseInputsRoleUnion1.BaseInputsRoleSystem1 is populated
	case components.BaseInputsRoleUnion1TypeBaseInputsRoleAssistant:
		// baseInputsRoleUnion1.BaseInputsRoleAssistant is populated
	case components.BaseInputsRoleUnion1TypeBaseInputsRoleDeveloper1:
		// baseInputsRoleUnion1.BaseInputsRoleDeveloper1 is populated
}
```
