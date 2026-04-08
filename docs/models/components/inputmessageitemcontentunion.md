# InputMessageItemContentUnion


## Supported Types

### InputText

```go
inputMessageItemContentUnion := components.CreateInputMessageItemContentUnionInputText(components.InputText{/* values here */})
```

### InputMessageItemContentInputImage

```go
inputMessageItemContentUnion := components.CreateInputMessageItemContentUnionInputImage(components.InputMessageItemContentInputImage{/* values here */})
```

### InputFile

```go
inputMessageItemContentUnion := components.CreateInputMessageItemContentUnionInputFile(components.InputFile{/* values here */})
```

### InputAudio

```go
inputMessageItemContentUnion := components.CreateInputMessageItemContentUnionInputAudio(components.InputAudio{/* values here */})
```

### InputVideo

```go
inputMessageItemContentUnion := components.CreateInputMessageItemContentUnionInputVideo(components.InputVideo{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch inputMessageItemContentUnion.Type {
	case components.InputMessageItemContentUnionTypeInputText:
		// inputMessageItemContentUnion.InputText is populated
	case components.InputMessageItemContentUnionTypeInputImage:
		// inputMessageItemContentUnion.InputMessageItemContentInputImage is populated
	case components.InputMessageItemContentUnionTypeInputFile:
		// inputMessageItemContentUnion.InputFile is populated
	case components.InputMessageItemContentUnionTypeInputAudio:
		// inputMessageItemContentUnion.InputAudio is populated
	case components.InputMessageItemContentUnionTypeInputVideo:
		// inputMessageItemContentUnion.InputVideo is populated
}
```
