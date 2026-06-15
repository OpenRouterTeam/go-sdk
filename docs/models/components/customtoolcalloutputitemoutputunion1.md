# CustomToolCallOutputItemOutputUnion1


## Supported Types

### InputText

```go
customToolCallOutputItemOutputUnion1 := components.CreateCustomToolCallOutputItemOutputUnion1InputText(components.InputText{/* values here */})
```

### CustomToolCallOutputItemOutputInputImage

```go
customToolCallOutputItemOutputUnion1 := components.CreateCustomToolCallOutputItemOutputUnion1InputImage(components.CustomToolCallOutputItemOutputInputImage{/* values here */})
```

### InputFile

```go
customToolCallOutputItemOutputUnion1 := components.CreateCustomToolCallOutputItemOutputUnion1InputFile(components.InputFile{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch customToolCallOutputItemOutputUnion1.Type {
	case components.CustomToolCallOutputItemOutputUnion1TypeInputText:
		// customToolCallOutputItemOutputUnion1.InputText is populated
	case components.CustomToolCallOutputItemOutputUnion1TypeInputImage:
		// customToolCallOutputItemOutputUnion1.CustomToolCallOutputItemOutputInputImage is populated
	case components.CustomToolCallOutputItemOutputUnion1TypeInputFile:
		// customToolCallOutputItemOutputUnion1.InputFile is populated
}
```
