# ImageGenerationServerToolConfigUnion


## Supported Types

### 

```go
imageGenerationServerToolConfigUnion := components.CreateImageGenerationServerToolConfigUnionStr(string{/* values here */})
```

### 

```go
imageGenerationServerToolConfigUnion := components.CreateImageGenerationServerToolConfigUnionNumber(float64{/* values here */})
```

### 

```go
imageGenerationServerToolConfigUnion := components.CreateImageGenerationServerToolConfigUnionArrayOfAny([]any{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch imageGenerationServerToolConfigUnion.Type {
	case components.ImageGenerationServerToolConfigUnionTypeStr:
		// imageGenerationServerToolConfigUnion.Str is populated
	case components.ImageGenerationServerToolConfigUnionTypeNumber:
		// imageGenerationServerToolConfigUnion.Number is populated
	case components.ImageGenerationServerToolConfigUnionTypeArrayOfAny:
		// imageGenerationServerToolConfigUnion.ArrayOfAny is populated
}
```
