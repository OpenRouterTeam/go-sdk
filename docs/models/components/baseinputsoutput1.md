# BaseInputsOutput1


## Supported Types

### InputText

```go
baseInputsOutput1 := components.CreateBaseInputsOutput1InputText(components.InputText{/* values here */})
```

### InputImage

```go
baseInputsOutput1 := components.CreateBaseInputsOutput1InputImage(components.InputImage{/* values here */})
```

### InputFile

```go
baseInputsOutput1 := components.CreateBaseInputsOutput1InputFile(components.InputFile{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch baseInputsOutput1.Type {
	case components.BaseInputsOutput1TypeInputText:
		// baseInputsOutput1.InputText is populated
	case components.BaseInputsOutput1TypeInputImage:
		// baseInputsOutput1.InputImage is populated
	case components.BaseInputsOutput1TypeInputFile:
		// baseInputsOutput1.InputFile is populated
}
```
