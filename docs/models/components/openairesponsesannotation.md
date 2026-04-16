# OpenAIResponsesAnnotation


## Supported Types

### FileCitation

```go
openAIResponsesAnnotation := components.CreateOpenAIResponsesAnnotationFileCitation(components.FileCitation{/* values here */})
```

### URLCitation

```go
openAIResponsesAnnotation := components.CreateOpenAIResponsesAnnotationURLCitation(components.URLCitation{/* values here */})
```

### FilePath

```go
openAIResponsesAnnotation := components.CreateOpenAIResponsesAnnotationFilePath(components.FilePath{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch openAIResponsesAnnotation.Type {
	case components.OpenAIResponsesAnnotationTypeFileCitation:
		// openAIResponsesAnnotation.FileCitation is populated
	case components.OpenAIResponsesAnnotationTypeURLCitation:
		// openAIResponsesAnnotation.URLCitation is populated
	case components.OpenAIResponsesAnnotationTypeFilePath:
		// openAIResponsesAnnotation.FilePath is populated
	default:
		// Unknown type - use openAIResponsesAnnotation.GetUnknownRaw() for raw JSON
}
```
