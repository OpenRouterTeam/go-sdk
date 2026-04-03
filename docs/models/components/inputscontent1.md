# InputsContent1


## Supported Types

### ResponseOutputText

```go
inputsContent1 := components.CreateInputsContent1OutputText(components.ResponseOutputText{/* values here */})
```

### OpenAIResponsesRefusalContent

```go
inputsContent1 := components.CreateInputsContent1Refusal(components.OpenAIResponsesRefusalContent{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch inputsContent1.Type {
	case components.InputsContent1TypeOutputText:
		// inputsContent1.ResponseOutputText is populated
	case components.InputsContent1TypeRefusal:
		// inputsContent1.OpenAIResponsesRefusalContent is populated
}
```
