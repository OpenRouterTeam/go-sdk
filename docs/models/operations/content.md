# Content


## Supported Types

### ContentText

```go
content := operations.CreateContentText(operations.ContentText{/* values here */})
```

### ContentImageURL

```go
content := operations.CreateContentImageURL(operations.ContentImageURL{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch content.Type {
	case operations.ContentTypeText:
		// content.ContentText is populated
	case operations.ContentTypeImageURL:
		// content.ContentImageURL is populated
}
```
