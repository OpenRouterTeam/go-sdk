# BaseInputsUnion1


## Supported Types

### BaseInputsMessage1

```go
baseInputsUnion1 := components.CreateBaseInputsUnion1BaseInputsMessage1(components.BaseInputsMessage1{/* values here */})
```

### BaseInputsMessage2

```go
baseInputsUnion1 := components.CreateBaseInputsUnion1BaseInputsMessage2(components.BaseInputsMessage2{/* values here */})
```

### BaseInputsFunctionCallOutput

```go
baseInputsUnion1 := components.CreateBaseInputsUnion1BaseInputsFunctionCallOutput(components.BaseInputsFunctionCallOutput{/* values here */})
```

### BaseInputsFunctionCall

```go
baseInputsUnion1 := components.CreateBaseInputsUnion1BaseInputsFunctionCall(components.BaseInputsFunctionCall{/* values here */})
```

### OutputItemImageGenerationCall

```go
baseInputsUnion1 := components.CreateBaseInputsUnion1OutputItemImageGenerationCall(components.OutputItemImageGenerationCall{/* values here */})
```

### OutputMessage

```go
baseInputsUnion1 := components.CreateBaseInputsUnion1OutputMessage(components.OutputMessage{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch baseInputsUnion1.Type {
	case components.BaseInputsUnion1TypeBaseInputsMessage1:
		// baseInputsUnion1.BaseInputsMessage1 is populated
	case components.BaseInputsUnion1TypeBaseInputsMessage2:
		// baseInputsUnion1.BaseInputsMessage2 is populated
	case components.BaseInputsUnion1TypeBaseInputsFunctionCallOutput:
		// baseInputsUnion1.BaseInputsFunctionCallOutput is populated
	case components.BaseInputsUnion1TypeBaseInputsFunctionCall:
		// baseInputsUnion1.BaseInputsFunctionCall is populated
	case components.BaseInputsUnion1TypeOutputItemImageGenerationCall:
		// baseInputsUnion1.OutputItemImageGenerationCall is populated
	case components.BaseInputsUnion1TypeOutputMessage:
		// baseInputsUnion1.OutputMessage is populated
}
```
