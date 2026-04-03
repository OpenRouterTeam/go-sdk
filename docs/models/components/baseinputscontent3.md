# BaseInputsContent3


## Supported Types

### InputText

```go
baseInputsContent3 := components.CreateBaseInputsContent3InputText(components.InputText{/* values here */})
```

### InputImage

```go
baseInputsContent3 := components.CreateBaseInputsContent3InputImage(components.InputImage{/* values here */})
```

### InputFile

```go
baseInputsContent3 := components.CreateBaseInputsContent3InputFile(components.InputFile{/* values here */})
```

### InputAudio

```go
baseInputsContent3 := components.CreateBaseInputsContent3InputAudio(components.InputAudio{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch baseInputsContent3.Type {
	case components.BaseInputsContent3TypeInputText:
		// baseInputsContent3.InputText is populated
	case components.BaseInputsContent3TypeInputImage:
		// baseInputsContent3.InputImage is populated
	case components.BaseInputsContent3TypeInputFile:
		// baseInputsContent3.InputFile is populated
	case components.BaseInputsContent3TypeInputAudio:
		// baseInputsContent3.InputAudio is populated
}
```
