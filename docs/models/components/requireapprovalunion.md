# RequireApprovalUnion


## Supported Types

### RequireApproval

```go
requireApprovalUnion := components.CreateRequireApprovalUnionRequireApproval(components.RequireApproval{/* values here */})
```

### RequireApprovalAlways

```go
requireApprovalUnion := components.CreateRequireApprovalUnionRequireApprovalAlways(components.RequireApprovalAlways{/* values here */})
```

### RequireApprovalNever

```go
requireApprovalUnion := components.CreateRequireApprovalUnionRequireApprovalNever(components.RequireApprovalNever{/* values here */})
```

### 

```go
requireApprovalUnion := components.CreateRequireApprovalUnionAny(any{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch requireApprovalUnion.Type {
	case components.RequireApprovalUnionTypeRequireApproval:
		// requireApprovalUnion.RequireApproval is populated
	case components.RequireApprovalUnionTypeRequireApprovalAlways:
		// requireApprovalUnion.RequireApprovalAlways is populated
	case components.RequireApprovalUnionTypeRequireApprovalNever:
		// requireApprovalUnion.RequireApprovalNever is populated
	case components.RequireApprovalUnionTypeAny:
		// requireApprovalUnion.Any is populated
}
```
