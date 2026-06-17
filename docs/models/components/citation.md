# Citation


## Supported Types

### AnthropicCitationCharLocationParam

```go
citation := components.CreateCitationCharLocation(components.AnthropicCitationCharLocationParam{/* values here */})
```

### AnthropicCitationContentBlockLocationParam

```go
citation := components.CreateCitationContentBlockLocation(components.AnthropicCitationContentBlockLocationParam{/* values here */})
```

### AnthropicCitationPageLocationParam

```go
citation := components.CreateCitationPageLocation(components.AnthropicCitationPageLocationParam{/* values here */})
```

### AnthropicCitationSearchResultLocation

```go
citation := components.CreateCitationSearchResultLocation(components.AnthropicCitationSearchResultLocation{/* values here */})
```

### AnthropicCitationWebSearchResultLocation

```go
citation := components.CreateCitationWebSearchResultLocation(components.AnthropicCitationWebSearchResultLocation{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch citation.Type {
	case components.CitationTypeCharLocation:
		// citation.AnthropicCitationCharLocationParam is populated
	case components.CitationTypeContentBlockLocation:
		// citation.AnthropicCitationContentBlockLocationParam is populated
	case components.CitationTypePageLocation:
		// citation.AnthropicCitationPageLocationParam is populated
	case components.CitationTypeSearchResultLocation:
		// citation.AnthropicCitationSearchResultLocation is populated
	case components.CitationTypeWebSearchResultLocation:
		// citation.AnthropicCitationWebSearchResultLocation is populated
}
```
