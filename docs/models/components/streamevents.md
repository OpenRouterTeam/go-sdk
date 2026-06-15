# StreamEvents

Union of all possible event types emitted during response streaming


## Supported Types

### ErrorEvent

```go
streamEvents := components.CreateStreamEventsError(components.ErrorEvent{/* values here */})
```

### ApplyPatchCallOperationDiffDeltaEvent

```go
streamEvents := components.CreateStreamEventsResponseApplyPatchCallOperationDiffDelta(components.ApplyPatchCallOperationDiffDeltaEvent{/* values here */})
```

### ApplyPatchCallOperationDiffDoneEvent

```go
streamEvents := components.CreateStreamEventsResponseApplyPatchCallOperationDiffDone(components.ApplyPatchCallOperationDiffDoneEvent{/* values here */})
```

### StreamEventsResponseCompleted

```go
streamEvents := components.CreateStreamEventsResponseCompleted(components.StreamEventsResponseCompleted{/* values here */})
```

### ContentPartAddedEvent

```go
streamEvents := components.CreateStreamEventsResponseContentPartAdded(components.ContentPartAddedEvent{/* values here */})
```

### ContentPartDoneEvent

```go
streamEvents := components.CreateStreamEventsResponseContentPartDone(components.ContentPartDoneEvent{/* values here */})
```

### OpenResponsesCreatedEvent

```go
streamEvents := components.CreateStreamEventsResponseCreated(components.OpenResponsesCreatedEvent{/* values here */})
```

### CustomToolCallInputDeltaEvent

```go
streamEvents := components.CreateStreamEventsResponseCustomToolCallInputDelta(components.CustomToolCallInputDeltaEvent{/* values here */})
```

### CustomToolCallInputDoneEvent

```go
streamEvents := components.CreateStreamEventsResponseCustomToolCallInputDone(components.CustomToolCallInputDoneEvent{/* values here */})
```

### StreamEventsResponseFailed

```go
streamEvents := components.CreateStreamEventsResponseFailed(components.StreamEventsResponseFailed{/* values here */})
```

### FunctionCallArgsDeltaEvent

```go
streamEvents := components.CreateStreamEventsResponseFunctionCallArgumentsDelta(components.FunctionCallArgsDeltaEvent{/* values here */})
```

### FunctionCallArgsDoneEvent

```go
streamEvents := components.CreateStreamEventsResponseFunctionCallArgumentsDone(components.FunctionCallArgsDoneEvent{/* values here */})
```

### FusionCallAnalysisCompletedEvent

```go
streamEvents := components.CreateStreamEventsResponseFusionCallAnalysisCompleted(components.FusionCallAnalysisCompletedEvent{/* values here */})
```

### FusionCallAnalysisInProgressEvent

```go
streamEvents := components.CreateStreamEventsResponseFusionCallAnalysisInProgress(components.FusionCallAnalysisInProgressEvent{/* values here */})
```

### FusionCallCompletedEvent

```go
streamEvents := components.CreateStreamEventsResponseFusionCallCompleted(components.FusionCallCompletedEvent{/* values here */})
```

### FusionCallInProgressEvent

```go
streamEvents := components.CreateStreamEventsResponseFusionCallInProgress(components.FusionCallInProgressEvent{/* values here */})
```

### FusionCallPanelAddedEvent

```go
streamEvents := components.CreateStreamEventsResponseFusionCallPanelAdded(components.FusionCallPanelAddedEvent{/* values here */})
```

### FusionCallPanelCompletedEvent

```go
streamEvents := components.CreateStreamEventsResponseFusionCallPanelCompleted(components.FusionCallPanelCompletedEvent{/* values here */})
```

### FusionCallPanelDeltaEvent

```go
streamEvents := components.CreateStreamEventsResponseFusionCallPanelDelta(components.FusionCallPanelDeltaEvent{/* values here */})
```

### FusionCallPanelFailedEvent

```go
streamEvents := components.CreateStreamEventsResponseFusionCallPanelFailed(components.FusionCallPanelFailedEvent{/* values here */})
```

### FusionCallPanelReasoningDeltaEvent

```go
streamEvents := components.CreateStreamEventsResponseFusionCallPanelReasoningDelta(components.FusionCallPanelReasoningDeltaEvent{/* values here */})
```

### ImageGenCallCompletedEvent

```go
streamEvents := components.CreateStreamEventsResponseImageGenerationCallCompleted(components.ImageGenCallCompletedEvent{/* values here */})
```

### ImageGenCallGeneratingEvent

```go
streamEvents := components.CreateStreamEventsResponseImageGenerationCallGenerating(components.ImageGenCallGeneratingEvent{/* values here */})
```

### ImageGenCallInProgressEvent

```go
streamEvents := components.CreateStreamEventsResponseImageGenerationCallInProgress(components.ImageGenCallInProgressEvent{/* values here */})
```

### ImageGenCallPartialImageEvent

```go
streamEvents := components.CreateStreamEventsResponseImageGenerationCallPartialImage(components.ImageGenCallPartialImageEvent{/* values here */})
```

### OpenResponsesInProgressEvent

```go
streamEvents := components.CreateStreamEventsResponseInProgress(components.OpenResponsesInProgressEvent{/* values here */})
```

### StreamEventsResponseIncomplete

```go
streamEvents := components.CreateStreamEventsResponseIncomplete(components.StreamEventsResponseIncomplete{/* values here */})
```

### StreamEventsResponseOutputItemAdded

```go
streamEvents := components.CreateStreamEventsResponseOutputItemAdded(components.StreamEventsResponseOutputItemAdded{/* values here */})
```

### StreamEventsResponseOutputItemDone

```go
streamEvents := components.CreateStreamEventsResponseOutputItemDone(components.StreamEventsResponseOutputItemDone{/* values here */})
```

### AnnotationAddedEvent

```go
streamEvents := components.CreateStreamEventsResponseOutputTextAnnotationAdded(components.AnnotationAddedEvent{/* values here */})
```

### TextDeltaEvent

```go
streamEvents := components.CreateStreamEventsResponseOutputTextDelta(components.TextDeltaEvent{/* values here */})
```

### TextDoneEvent

```go
streamEvents := components.CreateStreamEventsResponseOutputTextDone(components.TextDoneEvent{/* values here */})
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

### ReasoningDeltaEvent

```go
streamEvents := components.CreateStreamEventsResponseReasoningTextDelta(components.ReasoningDeltaEvent{/* values here */})
```

### ReasoningDoneEvent

```go
streamEvents := components.CreateStreamEventsResponseReasoningTextDone(components.ReasoningDoneEvent{/* values here */})
```

### RefusalDeltaEvent

```go
streamEvents := components.CreateStreamEventsResponseRefusalDelta(components.RefusalDeltaEvent{/* values here */})
```

### RefusalDoneEvent

```go
streamEvents := components.CreateStreamEventsResponseRefusalDone(components.RefusalDoneEvent{/* values here */})
```

### WebSearchCallCompletedEvent

```go
streamEvents := components.CreateStreamEventsResponseWebSearchCallCompleted(components.WebSearchCallCompletedEvent{/* values here */})
```

### WebSearchCallInProgressEvent

```go
streamEvents := components.CreateStreamEventsResponseWebSearchCallInProgress(components.WebSearchCallInProgressEvent{/* values here */})
```

### WebSearchCallSearchingEvent

```go
streamEvents := components.CreateStreamEventsResponseWebSearchCallSearching(components.WebSearchCallSearchingEvent{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch streamEvents.Type {
	case components.StreamEventsTypeError:
		// streamEvents.ErrorEvent is populated
	case components.StreamEventsTypeResponseApplyPatchCallOperationDiffDelta:
		// streamEvents.ApplyPatchCallOperationDiffDeltaEvent is populated
	case components.StreamEventsTypeResponseApplyPatchCallOperationDiffDone:
		// streamEvents.ApplyPatchCallOperationDiffDoneEvent is populated
	case components.StreamEventsTypeResponseCompleted:
		// streamEvents.StreamEventsResponseCompleted is populated
	case components.StreamEventsTypeResponseContentPartAdded:
		// streamEvents.ContentPartAddedEvent is populated
	case components.StreamEventsTypeResponseContentPartDone:
		// streamEvents.ContentPartDoneEvent is populated
	case components.StreamEventsTypeResponseCreated:
		// streamEvents.OpenResponsesCreatedEvent is populated
	case components.StreamEventsTypeResponseCustomToolCallInputDelta:
		// streamEvents.CustomToolCallInputDeltaEvent is populated
	case components.StreamEventsTypeResponseCustomToolCallInputDone:
		// streamEvents.CustomToolCallInputDoneEvent is populated
	case components.StreamEventsTypeResponseFailed:
		// streamEvents.StreamEventsResponseFailed is populated
	case components.StreamEventsTypeResponseFunctionCallArgumentsDelta:
		// streamEvents.FunctionCallArgsDeltaEvent is populated
	case components.StreamEventsTypeResponseFunctionCallArgumentsDone:
		// streamEvents.FunctionCallArgsDoneEvent is populated
	case components.StreamEventsTypeResponseFusionCallAnalysisCompleted:
		// streamEvents.FusionCallAnalysisCompletedEvent is populated
	case components.StreamEventsTypeResponseFusionCallAnalysisInProgress:
		// streamEvents.FusionCallAnalysisInProgressEvent is populated
	case components.StreamEventsTypeResponseFusionCallCompleted:
		// streamEvents.FusionCallCompletedEvent is populated
	case components.StreamEventsTypeResponseFusionCallInProgress:
		// streamEvents.FusionCallInProgressEvent is populated
	case components.StreamEventsTypeResponseFusionCallPanelAdded:
		// streamEvents.FusionCallPanelAddedEvent is populated
	case components.StreamEventsTypeResponseFusionCallPanelCompleted:
		// streamEvents.FusionCallPanelCompletedEvent is populated
	case components.StreamEventsTypeResponseFusionCallPanelDelta:
		// streamEvents.FusionCallPanelDeltaEvent is populated
	case components.StreamEventsTypeResponseFusionCallPanelFailed:
		// streamEvents.FusionCallPanelFailedEvent is populated
	case components.StreamEventsTypeResponseFusionCallPanelReasoningDelta:
		// streamEvents.FusionCallPanelReasoningDeltaEvent is populated
	case components.StreamEventsTypeResponseImageGenerationCallCompleted:
		// streamEvents.ImageGenCallCompletedEvent is populated
	case components.StreamEventsTypeResponseImageGenerationCallGenerating:
		// streamEvents.ImageGenCallGeneratingEvent is populated
	case components.StreamEventsTypeResponseImageGenerationCallInProgress:
		// streamEvents.ImageGenCallInProgressEvent is populated
	case components.StreamEventsTypeResponseImageGenerationCallPartialImage:
		// streamEvents.ImageGenCallPartialImageEvent is populated
	case components.StreamEventsTypeResponseInProgress:
		// streamEvents.OpenResponsesInProgressEvent is populated
	case components.StreamEventsTypeResponseIncomplete:
		// streamEvents.StreamEventsResponseIncomplete is populated
	case components.StreamEventsTypeResponseOutputItemAdded:
		// streamEvents.StreamEventsResponseOutputItemAdded is populated
	case components.StreamEventsTypeResponseOutputItemDone:
		// streamEvents.StreamEventsResponseOutputItemDone is populated
	case components.StreamEventsTypeResponseOutputTextAnnotationAdded:
		// streamEvents.AnnotationAddedEvent is populated
	case components.StreamEventsTypeResponseOutputTextDelta:
		// streamEvents.TextDeltaEvent is populated
	case components.StreamEventsTypeResponseOutputTextDone:
		// streamEvents.TextDoneEvent is populated
	case components.StreamEventsTypeResponseReasoningSummaryPartAdded:
		// streamEvents.ReasoningSummaryPartAddedEvent is populated
	case components.StreamEventsTypeResponseReasoningSummaryPartDone:
		// streamEvents.ReasoningSummaryPartDoneEvent is populated
	case components.StreamEventsTypeResponseReasoningSummaryTextDelta:
		// streamEvents.ReasoningSummaryTextDeltaEvent is populated
	case components.StreamEventsTypeResponseReasoningSummaryTextDone:
		// streamEvents.ReasoningSummaryTextDoneEvent is populated
	case components.StreamEventsTypeResponseReasoningTextDelta:
		// streamEvents.ReasoningDeltaEvent is populated
	case components.StreamEventsTypeResponseReasoningTextDone:
		// streamEvents.ReasoningDoneEvent is populated
	case components.StreamEventsTypeResponseRefusalDelta:
		// streamEvents.RefusalDeltaEvent is populated
	case components.StreamEventsTypeResponseRefusalDone:
		// streamEvents.RefusalDoneEvent is populated
	case components.StreamEventsTypeResponseWebSearchCallCompleted:
		// streamEvents.WebSearchCallCompletedEvent is populated
	case components.StreamEventsTypeResponseWebSearchCallInProgress:
		// streamEvents.WebSearchCallInProgressEvent is populated
	case components.StreamEventsTypeResponseWebSearchCallSearching:
		// streamEvents.WebSearchCallSearchingEvent is populated
	default:
		// Unknown type - use streamEvents.GetUnknownRaw() for raw JSON
}
```
