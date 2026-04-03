# EasyInputMessageContentUnion1


## Supported Types

### InputText

```go
easyInputMessageContentUnion1 := components.CreateEasyInputMessageContentUnion1InputText(components.InputText{/* values here */})
```

### EasyInputMessageContentInputImage

```go
easyInputMessageContentUnion1 := components.CreateEasyInputMessageContentUnion1InputImage(components.EasyInputMessageContentInputImage{/* values here */})
```

### InputFile

```go
easyInputMessageContentUnion1 := components.CreateEasyInputMessageContentUnion1InputFile(components.InputFile{/* values here */})
```

### InputAudio

```go
easyInputMessageContentUnion1 := components.CreateEasyInputMessageContentUnion1InputAudio(components.InputAudio{/* values here */})
```

### InputVideo

```go
easyInputMessageContentUnion1 := components.CreateEasyInputMessageContentUnion1InputVideo(components.InputVideo{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch easyInputMessageContentUnion1.Type {
	case components.EasyInputMessageContentUnion1TypeInputText:
		// easyInputMessageContentUnion1.InputText is populated
	case components.EasyInputMessageContentUnion1TypeInputImage:
		// easyInputMessageContentUnion1.EasyInputMessageContentInputImage is populated
	case components.EasyInputMessageContentUnion1TypeInputFile:
		// easyInputMessageContentUnion1.InputFile is populated
	case components.EasyInputMessageContentUnion1TypeInputAudio:
		// easyInputMessageContentUnion1.InputAudio is populated
	case components.EasyInputMessageContentUnion1TypeInputVideo:
		// easyInputMessageContentUnion1.InputVideo is populated
}
```
