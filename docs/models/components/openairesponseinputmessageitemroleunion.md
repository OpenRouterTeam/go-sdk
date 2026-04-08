# OpenAIResponseInputMessageItemRoleUnion


## Supported Types

### OpenAIResponseInputMessageItemRoleUser

```go
openAIResponseInputMessageItemRoleUnion := components.CreateOpenAIResponseInputMessageItemRoleUnionOpenAIResponseInputMessageItemRoleUser(components.OpenAIResponseInputMessageItemRoleUser{/* values here */})
```

### OpenAIResponseInputMessageItemRoleSystem

```go
openAIResponseInputMessageItemRoleUnion := components.CreateOpenAIResponseInputMessageItemRoleUnionOpenAIResponseInputMessageItemRoleSystem(components.OpenAIResponseInputMessageItemRoleSystem{/* values here */})
```

### OpenAIResponseInputMessageItemRoleDeveloper

```go
openAIResponseInputMessageItemRoleUnion := components.CreateOpenAIResponseInputMessageItemRoleUnionOpenAIResponseInputMessageItemRoleDeveloper(components.OpenAIResponseInputMessageItemRoleDeveloper{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch openAIResponseInputMessageItemRoleUnion.Type {
	case components.OpenAIResponseInputMessageItemRoleUnionTypeOpenAIResponseInputMessageItemRoleUser:
		// openAIResponseInputMessageItemRoleUnion.OpenAIResponseInputMessageItemRoleUser is populated
	case components.OpenAIResponseInputMessageItemRoleUnionTypeOpenAIResponseInputMessageItemRoleSystem:
		// openAIResponseInputMessageItemRoleUnion.OpenAIResponseInputMessageItemRoleSystem is populated
	case components.OpenAIResponseInputMessageItemRoleUnionTypeOpenAIResponseInputMessageItemRoleDeveloper:
		// openAIResponseInputMessageItemRoleUnion.OpenAIResponseInputMessageItemRoleDeveloper is populated
}
```
