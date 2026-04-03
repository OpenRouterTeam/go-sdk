# FunctionCallOutputItemOutputUnion1


## Supported Types

### InputText

```go
functionCallOutputItemOutputUnion1 := components.CreateFunctionCallOutputItemOutputUnion1InputText(components.InputText{/* values here */})
```

### OutputInputImage

```go
functionCallOutputItemOutputUnion1 := components.CreateFunctionCallOutputItemOutputUnion1InputImage(components.OutputInputImage{/* values here */})
```

### InputFile

```go
functionCallOutputItemOutputUnion1 := components.CreateFunctionCallOutputItemOutputUnion1InputFile(components.InputFile{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch functionCallOutputItemOutputUnion1.Type {
	case components.FunctionCallOutputItemOutputUnion1TypeInputText:
		// functionCallOutputItemOutputUnion1.InputText is populated
	case components.FunctionCallOutputItemOutputUnion1TypeInputImage:
		// functionCallOutputItemOutputUnion1.OutputInputImage is populated
	case components.FunctionCallOutputItemOutputUnion1TypeInputFile:
		// functionCallOutputItemOutputUnion1.InputFile is populated
}
```
