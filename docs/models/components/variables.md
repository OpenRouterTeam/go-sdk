# Variables


## Supported Types

### 

```go
variables := components.CreateVariablesStr(string{/* values here */})
```

### InputText

```go
variables := components.CreateVariablesInputText(components.InputText{/* values here */})
```

### InputImage

```go
variables := components.CreateVariablesInputImage(components.InputImage{/* values here */})
```

### InputFile

```go
variables := components.CreateVariablesInputFile(components.InputFile{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch variables.Type {
	case components.VariablesTypeStr:
		// variables.Str is populated
	case components.VariablesTypeInputText:
		// variables.InputText is populated
	case components.VariablesTypeInputImage:
		// variables.InputImage is populated
	case components.VariablesTypeInputFile:
		// variables.InputFile is populated
}
```
