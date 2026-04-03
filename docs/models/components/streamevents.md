# StreamEvents

Union of all possible event types emitted during response streaming


## Supported Types

### StreamEventsResponseCreated

```go
streamEvents := components.CreateStreamEventsResponseCreated(components.StreamEventsResponseCreated{/* values here */})
```

### StreamEventsResponseInProgress

```go
streamEvents := components.CreateStreamEventsResponseInProgress(components.StreamEventsResponseInProgress{/* values here */})
```

### StreamEventsResponseCompleted

```go
streamEvents := components.CreateStreamEventsResponseCompleted(components.StreamEventsResponseCompleted{/* values here */})
```

### StreamEventsResponseIncomplete

```go
streamEvents := components.CreateStreamEventsResponseIncomplete(components.StreamEventsResponseIncomplete{/* values here */})
```

### StreamEventsResponseFailed

```go
streamEvents := components.CreateStreamEventsResponseFailed(components.StreamEventsResponseFailed{/* values here */})
```

### ErrorEvent

```go
streamEvents := components.CreateStreamEventsError(components.ErrorEvent{/* values here */})
```

### StreamEventsResponseOutputItemAdded

```go
streamEvents := components.CreateStreamEventsResponseOutputItemAdded(components.StreamEventsResponseOutputItemAdded{/* values here */})
```

### StreamEventsResponseOutputItemDone

```go
streamEvents := components.CreateStreamEventsResponseOutputItemDone(components.StreamEventsResponseOutputItemDone{/* values here */})
```

### ContentPartAddedEvent

```go
streamEvents := components.CreateStreamEventsResponseContentPartAdded(components.ContentPartAddedEvent{/* values here */})
```

### ContentPartDoneEvent

```go
streamEvents := components.CreateStreamEventsResponseContentPartDone(components.ContentPartDoneEvent{/* values here */})
```

### TextDeltaEvent

```go
streamEvents := components.CreateStreamEventsResponseOutputTextDelta(components.TextDeltaEvent{/* values here */})
```

### TextDoneEvent

```go
streamEvents := components.CreateStreamEventsResponseOutputTextDone(components.TextDoneEvent{/* values here */})
```

### RefusalDeltaEvent

```go
streamEvents := components.CreateStreamEventsResponseRefusalDelta(components.RefusalDeltaEvent{/* values here */})
```

### RefusalDoneEvent

```go
streamEvents := components.CreateStreamEventsResponseRefusalDone(components.RefusalDoneEvent{/* values here */})
```

### AnnotationAddedEvent

```go
streamEvents := components.CreateStreamEventsResponseOutputTextAnnotationAdded(components.AnnotationAddedEvent{/* values here */})
```

### FunctionCallArgsDeltaEvent

```go
streamEvents := components.CreateStreamEventsResponseFunctionCallArgumentsDelta(components.FunctionCallArgsDeltaEvent{/* values here */})
```

### FunctionCallArgsDoneEvent

```go
streamEvents := components.CreateStreamEventsResponseFunctionCallArgumentsDone(components.FunctionCallArgsDoneEvent{/* values here */})
```

### ReasoningDeltaEvent

```go
streamEvents := components.CreateStreamEventsResponseReasoningTextDelta(components.ReasoningDeltaEvent{/* values here */})
```

### ReasoningDoneEvent

```go
streamEvents := components.CreateStreamEventsResponseReasoningTextDone(components.ReasoningDoneEvent{/* values here */})
```

### ReasoningSummaryPartAddedEvent

```go
streamEvents := components.CreateStreamEventsResponseReasoningSummaryPartAdded(components.ReasoningSummaryPartAddedEvent{/* values here */})
```

### ReasoningSummaryPartDoneEvent

```go
streamEvents := components.CreateStreamEventsResponseReasoningSummaryPartDone(components.ReasoningSummaryPartDoneEvent{/* values here */})
```

### ReasoningSummaryTextDeltaEvent

```go
streamEvents := components.CreateStreamEventsResponseReasoningSummaryTextDelta(components.ReasoningSummaryTextDeltaEvent{/* values here */})
```

### ReasoningSummaryTextDoneEvent

```go
streamEvents := components.CreateStreamEventsResponseReasoningSummaryTextDone(components.ReasoningSummaryTextDoneEvent{/* values here */})
```

### ImageGenCallInProgressEvent

```go
streamEvents := components.CreateStreamEventsResponseImageGenerationCallInProgress(components.ImageGenCallInProgressEvent{/* values here */})
```

### ImageGenCallGeneratingEvent

```go
streamEvents := components.CreateStreamEventsResponseImageGenerationCallGenerating(components.ImageGenCallGeneratingEvent{/* values here */})
```

### ImageGenCallPartialImageEvent

```go
streamEvents := components.CreateStreamEventsResponseImageGenerationCallPartialImage(components.ImageGenCallPartialImageEvent{/* values here */})
```

### ImageGenCallCompletedEvent

```go
streamEvents := components.CreateStreamEventsResponseImageGenerationCallCompleted(components.ImageGenCallCompletedEvent{/* values here */})
```

### WebSearchCallInProgressEvent

```go
streamEvents := components.CreateStreamEventsResponseWebSearchCallInProgress(components.WebSearchCallInProgressEvent{/* values here */})
```

### WebSearchCallSearchingEvent

```go
streamEvents := components.CreateStreamEventsResponseWebSearchCallSearching(components.WebSearchCallSearchingEvent{/* values here */})
```

### WebSearchCallCompletedEvent

```go
streamEvents := components.CreateStreamEventsResponseWebSearchCallCompleted(components.WebSearchCallCompletedEvent{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch streamEvents.Type {
	case components.StreamEventsTypeResponseCreated:
		// streamEvents.StreamEventsResponseCreated is populated
	case components.StreamEventsTypeResponseInProgress:
		// streamEvents.StreamEventsResponseInProgress is populated
	case components.StreamEventsTypeResponseCompleted:
		// streamEvents.StreamEventsResponseCompleted is populated
	case components.StreamEventsTypeResponseIncomplete:
		// streamEvents.StreamEventsResponseIncomplete is populated
	case components.StreamEventsTypeResponseFailed:
		// streamEvents.StreamEventsResponseFailed is populated
	case components.StreamEventsTypeError:
		// streamEvents.ErrorEvent is populated
	case components.StreamEventsTypeResponseOutputItemAdded:
		// streamEvents.StreamEventsResponseOutputItemAdded is populated
	case components.StreamEventsTypeResponseOutputItemDone:
		// streamEvents.StreamEventsResponseOutputItemDone is populated
	case components.StreamEventsTypeResponseContentPartAdded:
		// streamEvents.ContentPartAddedEvent is populated
	case components.StreamEventsTypeResponseContentPartDone:
		// streamEvents.ContentPartDoneEvent is populated
	case components.StreamEventsTypeResponseOutputTextDelta:
		// streamEvents.TextDeltaEvent is populated
	case components.StreamEventsTypeResponseOutputTextDone:
		// streamEvents.TextDoneEvent is populated
	case components.StreamEventsTypeResponseRefusalDelta:
		// streamEvents.RefusalDeltaEvent is populated
	case components.StreamEventsTypeResponseRefusalDone:
		// streamEvents.RefusalDoneEvent is populated
	case components.StreamEventsTypeResponseOutputTextAnnotationAdded:
		// streamEvents.AnnotationAddedEvent is populated
	case components.StreamEventsTypeResponseFunctionCallArgumentsDelta:
		// streamEvents.FunctionCallArgsDeltaEvent is populated
	case components.StreamEventsTypeResponseFunctionCallArgumentsDone:
		// streamEvents.FunctionCallArgsDoneEvent is populated
	case components.StreamEventsTypeResponseReasoningTextDelta:
		// streamEvents.ReasoningDeltaEvent is populated
	case components.StreamEventsTypeResponseReasoningTextDone:
		// streamEvents.ReasoningDoneEvent is populated
	case components.StreamEventsTypeResponseReasoningSummaryPartAdded:
		// streamEvents.ReasoningSummaryPartAddedEvent is populated
	case components.StreamEventsTypeResponseReasoningSummaryPartDone:
		// streamEvents.ReasoningSummaryPartDoneEvent is populated
	case components.StreamEventsTypeResponseReasoningSummaryTextDelta:
		// streamEvents.ReasoningSummaryTextDeltaEvent is populated
	case components.StreamEventsTypeResponseReasoningSummaryTextDone:
		// streamEvents.ReasoningSummaryTextDoneEvent is populated
	case components.StreamEventsTypeResponseImageGenerationCallInProgress:
		// streamEvents.ImageGenCallInProgressEvent is populated
	case components.StreamEventsTypeResponseImageGenerationCallGenerating:
		// streamEvents.ImageGenCallGeneratingEvent is populated
	case components.StreamEventsTypeResponseImageGenerationCallPartialImage:
		// streamEvents.ImageGenCallPartialImageEvent is populated
	case components.StreamEventsTypeResponseImageGenerationCallCompleted:
		// streamEvents.ImageGenCallCompletedEvent is populated
	case components.StreamEventsTypeResponseWebSearchCallInProgress:
		// streamEvents.WebSearchCallInProgressEvent is populated
	case components.StreamEventsTypeResponseWebSearchCallSearching:
		// streamEvents.WebSearchCallSearchingEvent is populated
	case components.StreamEventsTypeResponseWebSearchCallCompleted:
		// streamEvents.WebSearchCallCompletedEvent is populated
}
```
