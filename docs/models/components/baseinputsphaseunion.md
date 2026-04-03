# BaseInputsPhaseUnion


## Supported Types

### BaseInputsPhaseCommentary

```go
baseInputsPhaseUnion := components.CreateBaseInputsPhaseUnionBaseInputsPhaseCommentary(components.BaseInputsPhaseCommentary{/* values here */})
```

### BaseInputsPhaseFinalAnswer

```go
baseInputsPhaseUnion := components.CreateBaseInputsPhaseUnionBaseInputsPhaseFinalAnswer(components.BaseInputsPhaseFinalAnswer{/* values here */})
```

### 

```go
baseInputsPhaseUnion := components.CreateBaseInputsPhaseUnionAny(any{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch baseInputsPhaseUnion.Type {
	case components.BaseInputsPhaseUnionTypeBaseInputsPhaseCommentary:
		// baseInputsPhaseUnion.BaseInputsPhaseCommentary is populated
	case components.BaseInputsPhaseUnionTypeBaseInputsPhaseFinalAnswer:
		// baseInputsPhaseUnion.BaseInputsPhaseFinalAnswer is populated
	case components.BaseInputsPhaseUnionTypeAny:
		// baseInputsPhaseUnion.Any is populated
}
```
