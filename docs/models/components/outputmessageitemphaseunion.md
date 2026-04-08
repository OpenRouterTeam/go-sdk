# OutputMessageItemPhaseUnion

The phase of an assistant message. Use `commentary` for an intermediate assistant message and `final_answer` for the final assistant message. For follow-up requests with models like `gpt-5.3-codex` and later, preserve and resend phase on all assistant messages. Omitting it can degrade performance. Not used for user messages.


## Supported Types

### OutputMessageItemPhaseCommentary

```go
outputMessageItemPhaseUnion := components.CreateOutputMessageItemPhaseUnionOutputMessageItemPhaseCommentary(components.OutputMessageItemPhaseCommentary{/* values here */})
```

### OutputMessageItemPhaseFinalAnswer

```go
outputMessageItemPhaseUnion := components.CreateOutputMessageItemPhaseUnionOutputMessageItemPhaseFinalAnswer(components.OutputMessageItemPhaseFinalAnswer{/* values here */})
```

### 

```go
outputMessageItemPhaseUnion := components.CreateOutputMessageItemPhaseUnionAny(any{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch outputMessageItemPhaseUnion.Type {
	case components.OutputMessageItemPhaseUnionTypeOutputMessageItemPhaseCommentary:
		// outputMessageItemPhaseUnion.OutputMessageItemPhaseCommentary is populated
	case components.OutputMessageItemPhaseUnionTypeOutputMessageItemPhaseFinalAnswer:
		// outputMessageItemPhaseUnion.OutputMessageItemPhaseFinalAnswer is populated
	case components.OutputMessageItemPhaseUnionTypeAny:
		// outputMessageItemPhaseUnion.Any is populated
}
```
