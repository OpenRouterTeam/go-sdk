# PDFParserEngine

The engine to use for parsing PDF files. "pdf-text" is deprecated and automatically redirected to "cloudflare-ai".


## Supported Types

### PDFParserEngineEnum

```go
pdfParserEngine := components.CreatePDFParserEnginePDFParserEngineEnum(components.PDFParserEngineEnum{/* values here */})
```

### PDFParserEnginePDFText

```go
pdfParserEngine := components.CreatePDFParserEnginePDFParserEnginePDFText(components.PDFParserEnginePDFText{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch pdfParserEngine.Type {
	case components.PDFParserEngineTypePDFParserEngineEnum:
		// pdfParserEngine.PDFParserEngineEnum is populated
	case components.PDFParserEngineTypePDFParserEnginePDFText:
		// pdfParserEngine.PDFParserEnginePDFText is populated
}
```
