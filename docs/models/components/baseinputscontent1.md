# BaseInputsContent1


## Supported Types

### InputText

```go
baseInputsContent1 := components.CreateBaseInputsContent1InputText(components.InputText{/* values here */})
```

### InputImage

```go
baseInputsContent1 := components.CreateBaseInputsContent1InputImage(components.InputImage{/* values here */})
```

### InputFile

```go
baseInputsContent1 := components.CreateBaseInputsContent1InputFile(components.InputFile{/* values here */})
```

### InputAudio

```go
baseInputsContent1 := components.CreateBaseInputsContent1InputAudio(components.InputAudio{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch baseInputsContent1.Type {
	case components.BaseInputsContent1TypeInputText:
		// baseInputsContent1.InputText is populated
	case components.BaseInputsContent1TypeInputImage:
		// baseInputsContent1.InputImage is populated
	case components.BaseInputsContent1TypeInputFile:
		// baseInputsContent1.InputFile is populated
	case components.BaseInputsContent1TypeInputAudio:
		// baseInputsContent1.InputAudio is populated
}
```
