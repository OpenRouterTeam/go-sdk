# Value1

Filter value (scalar or array depending on operator). Several dimensions are enriched in responses (returned as human-readable labels), but filters must use the underlying ID: `api_key_id` — numeric ID (from generation metadata) or key hash (64-char hex from GET /api/v1/keys, resolved server-side); `user` — Clerk user ID (e.g. "user_abc123"), not the display name; `workspace` — workspace UUID, not the workspace name; `app` — numeric app ID, not the app title; `model` — permaslug (e.g. "openai/gpt-4o"), not the display name. Other dimensions (provider, origin, country, etc.) are not enriched and accept the value as returned.


## Supported Types

### 

```go
value1 := operations.CreateValue1Str(string{/* values here */})
```

### 

```go
value1 := operations.CreateValue1Number(float64{/* values here */})
```

### 

```go
value1 := operations.CreateValue1ArrayOfValue2([]operations.Value2{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch value1.Type {
	case operations.Value1TypeStr:
		// value1.Str is populated
	case operations.Value1TypeNumber:
		// value1.Number is populated
	case operations.Value1TypeArrayOfValue2:
		// value1.ArrayOfValue2 is populated
}
```
