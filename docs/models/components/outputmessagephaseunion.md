# OutputMessagePhaseUnion

The phase of an assistant message. Use `commentary` for an intermediate assistant message and `final_answer` for the final assistant message. For follow-up requests with models like `gpt-5.3-codex` and later, preserve and resend phase on all assistant messages. Omitting it can degrade performance. Not used for user messages.


## Supported Types

### OutputMessagePhaseCommentary

```go
outputMessagePhaseUnion := components.CreateOutputMessagePhaseUnionOutputMessagePhaseCommentary(components.OutputMessagePhaseCommentary{/* values here */})
```

### OutputMessagePhaseFinalAnswer

```go
outputMessagePhaseUnion := components.CreateOutputMessagePhaseUnionOutputMessagePhaseFinalAnswer(components.OutputMessagePhaseFinalAnswer{/* values here */})
```

### 

```go
outputMessagePhaseUnion := components.CreateOutputMessagePhaseUnionAny(any{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch outputMessagePhaseUnion.Type {
	case components.OutputMessagePhaseUnionTypeOutputMessagePhaseCommentary:
		// outputMessagePhaseUnion.OutputMessagePhaseCommentary is populated
	case components.OutputMessagePhaseUnionTypeOutputMessagePhaseFinalAnswer:
		// outputMessagePhaseUnion.OutputMessagePhaseFinalAnswer is populated
	case components.OutputMessagePhaseUnionTypeAny:
		// outputMessagePhaseUnion.Any is populated
}
```
