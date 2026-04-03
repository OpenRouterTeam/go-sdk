# InputsPhaseUnion

The phase of an assistant message. Use `commentary` for an intermediate assistant message and `final_answer` for the final assistant message. For follow-up requests with models like `gpt-5.3-codex` and later, preserve and resend phase on all assistant messages. Omitting it can degrade performance. Not used for user messages.


## Supported Types

### InputsPhaseCommentary

```go
inputsPhaseUnion := components.CreateInputsPhaseUnionInputsPhaseCommentary(components.InputsPhaseCommentary{/* values here */})
```

### InputsPhaseFinalAnswer

```go
inputsPhaseUnion := components.CreateInputsPhaseUnionInputsPhaseFinalAnswer(components.InputsPhaseFinalAnswer{/* values here */})
```

### 

```go
inputsPhaseUnion := components.CreateInputsPhaseUnionAny(any{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch inputsPhaseUnion.Type {
	case components.InputsPhaseUnionTypeInputsPhaseCommentary:
		// inputsPhaseUnion.InputsPhaseCommentary is populated
	case components.InputsPhaseUnionTypeInputsPhaseFinalAnswer:
		// inputsPhaseUnion.InputsPhaseFinalAnswer is populated
	case components.InputsPhaseUnionTypeAny:
		// inputsPhaseUnion.Any is populated
}
```
