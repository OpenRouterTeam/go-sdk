# OpenAIResponseCustomToolCallOutputOutput2


## Supported Types

### 

```go
openAIResponseCustomToolCallOutputOutput2 := components.CreateOpenAIResponseCustomToolCallOutputOutput2Str(string{/* values here */})
```

### 

```go
openAIResponseCustomToolCallOutputOutput2 := components.CreateOpenAIResponseCustomToolCallOutputOutput2ArrayOfOpenAIResponseCustomToolCallOutputOutput1([]components.OpenAIResponseCustomToolCallOutputOutput1{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch openAIResponseCustomToolCallOutputOutput2.Type {
	case components.OpenAIResponseCustomToolCallOutputOutput2TypeStr:
		// openAIResponseCustomToolCallOutputOutput2.Str is populated
	case components.OpenAIResponseCustomToolCallOutputOutput2TypeArrayOfOpenAIResponseCustomToolCallOutputOutput1:
		// openAIResponseCustomToolCallOutputOutput2.ArrayOfOpenAIResponseCustomToolCallOutputOutput1 is populated
}
```
