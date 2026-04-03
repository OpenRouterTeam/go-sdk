# EasyInputMessagePhaseUnion

The phase of an assistant message. Use `commentary` for an intermediate assistant message and `final_answer` for the final assistant message. For follow-up requests with models like `gpt-5.3-codex` and later, preserve and resend phase on all assistant messages. Omitting it can degrade performance. Not used for user messages.


## Supported Types

### EasyInputMessagePhaseCommentary

```go
easyInputMessagePhaseUnion := components.CreateEasyInputMessagePhaseUnionEasyInputMessagePhaseCommentary(components.EasyInputMessagePhaseCommentary{/* values here */})
```

### EasyInputMessagePhaseFinalAnswer

```go
easyInputMessagePhaseUnion := components.CreateEasyInputMessagePhaseUnionEasyInputMessagePhaseFinalAnswer(components.EasyInputMessagePhaseFinalAnswer{/* values here */})
```

### 

```go
easyInputMessagePhaseUnion := components.CreateEasyInputMessagePhaseUnionAny(any{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch easyInputMessagePhaseUnion.Type {
	case components.EasyInputMessagePhaseUnionTypeEasyInputMessagePhaseCommentary:
		// easyInputMessagePhaseUnion.EasyInputMessagePhaseCommentary is populated
	case components.EasyInputMessagePhaseUnionTypeEasyInputMessagePhaseFinalAnswer:
		// easyInputMessagePhaseUnion.EasyInputMessagePhaseFinalAnswer is populated
	case components.EasyInputMessagePhaseUnionTypeAny:
		// easyInputMessagePhaseUnion.Any is populated
}
```
