# OutputCodeInterpreterCallItemOutputUnion


## Supported Types

### OutputImage

```go
outputCodeInterpreterCallItemOutputUnion := components.CreateOutputCodeInterpreterCallItemOutputUnionImage(components.OutputImage{/* values here */})
```

### OutputLogs

```go
outputCodeInterpreterCallItemOutputUnion := components.CreateOutputCodeInterpreterCallItemOutputUnionLogs(components.OutputLogs{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch outputCodeInterpreterCallItemOutputUnion.Type {
	case components.OutputCodeInterpreterCallItemOutputUnionTypeImage:
		// outputCodeInterpreterCallItemOutputUnion.OutputImage is populated
	case components.OutputCodeInterpreterCallItemOutputUnionTypeLogs:
		// outputCodeInterpreterCallItemOutputUnion.OutputLogs is populated
	default:
		// Unknown type - use outputCodeInterpreterCallItemOutputUnion.GetUnknownRaw() for raw JSON
}
```
