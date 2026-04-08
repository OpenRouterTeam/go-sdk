# OpenAIResponseFunctionToolCallOutputOutput2


## Supported Types

### 

```go
openAIResponseFunctionToolCallOutputOutput2 := components.CreateOpenAIResponseFunctionToolCallOutputOutput2Str(string{/* values here */})
```

### 

```go
openAIResponseFunctionToolCallOutputOutput2 := components.CreateOpenAIResponseFunctionToolCallOutputOutput2ArrayOfOpenAIResponseFunctionToolCallOutputOutput1([]components.OpenAIResponseFunctionToolCallOutputOutput1{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch openAIResponseFunctionToolCallOutputOutput2.Type {
	case components.OpenAIResponseFunctionToolCallOutputOutput2TypeStr:
		// openAIResponseFunctionToolCallOutputOutput2.Str is populated
	case components.OpenAIResponseFunctionToolCallOutputOutput2TypeArrayOfOpenAIResponseFunctionToolCallOutputOutput1:
		// openAIResponseFunctionToolCallOutputOutput2.ArrayOfOpenAIResponseFunctionToolCallOutputOutput1 is populated
}
```
