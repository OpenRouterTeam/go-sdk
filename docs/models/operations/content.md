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

### ContentPartInputAudio

```go
content := operations.CreateContentInputAudio(components.ContentPartInputAudio{/* values here */})
```

### ContentPartInputVideo

```go
content := operations.CreateContentInputVideo(components.ContentPartInputVideo{/* values here */})
```

### ContentPartInputFile

```go
content := operations.CreateContentInputFile(components.ContentPartInputFile{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch content.Type {
	case operations.ContentTypeText:
		// content.ContentText is populated
	case operations.ContentTypeImageURL:
		// content.ContentImageURL is populated
	case operations.ContentTypeInputAudio:
		// content.ContentPartInputAudio is populated
	case operations.ContentTypeInputVideo:
		// content.ContentPartInputVideo is populated
	case operations.ContentTypeInputFile:
		// content.ContentPartInputFile is populated
}
```
