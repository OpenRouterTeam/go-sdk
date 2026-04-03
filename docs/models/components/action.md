# Action


## Supported Types

### ActionSearch

```go
action := components.CreateActionSearch(components.ActionSearch{/* values here */})
```

### ActionOpenPage

```go
action := components.CreateActionOpenPage(components.ActionOpenPage{/* values here */})
```

### ActionFindInPage

```go
action := components.CreateActionFindInPage(components.ActionFindInPage{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch action.Type {
	case components.ActionTypeSearch:
		// action.ActionSearch is populated
	case components.ActionTypeOpenPage:
		// action.ActionOpenPage is populated
	case components.ActionTypeFindInPage:
		// action.ActionFindInPage is populated
}
```
