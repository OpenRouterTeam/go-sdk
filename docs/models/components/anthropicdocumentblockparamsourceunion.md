# AnthropicDocumentBlockParamSourceUnion


## Supported Types

### AnthropicBase64PdfSource

```go
anthropicDocumentBlockParamSourceUnion := components.CreateAnthropicDocumentBlockParamSourceUnionBase64(components.AnthropicBase64PdfSource{/* values here */})
```

### AnthropicPlainTextSource

```go
anthropicDocumentBlockParamSourceUnion := components.CreateAnthropicDocumentBlockParamSourceUnionText(components.AnthropicPlainTextSource{/* values here */})
```

### SourceContent

```go
anthropicDocumentBlockParamSourceUnion := components.CreateAnthropicDocumentBlockParamSourceUnionContent(components.SourceContent{/* values here */})
```

### AnthropicURLPdfSource

```go
anthropicDocumentBlockParamSourceUnion := components.CreateAnthropicDocumentBlockParamSourceUnionURLObj(components.AnthropicURLPdfSource{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch anthropicDocumentBlockParamSourceUnion.Type {
	case components.AnthropicDocumentBlockParamSourceUnionTypeBase64:
		// anthropicDocumentBlockParamSourceUnion.AnthropicBase64PdfSource is populated
	case components.AnthropicDocumentBlockParamSourceUnionTypeText:
		// anthropicDocumentBlockParamSourceUnion.AnthropicPlainTextSource is populated
	case components.AnthropicDocumentBlockParamSourceUnionTypeContent:
		// anthropicDocumentBlockParamSourceUnion.SourceContent is populated
	case components.AnthropicDocumentBlockParamSourceUnionTypeURLObj:
		// anthropicDocumentBlockParamSourceUnion.AnthropicURLPdfSource is populated
}
```
