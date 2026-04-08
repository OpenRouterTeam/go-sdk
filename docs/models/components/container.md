# Container


## Supported Types

### 

```go
container := components.CreateContainerStr(string{/* values here */})
```

### ContainerAuto

```go
container := components.CreateContainerContainerAuto(components.ContainerAuto{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch container.Type {
	case components.ContainerUnionTypeStr:
		// container.Str is populated
	case components.ContainerUnionTypeContainerAuto:
		// container.ContainerAuto is populated
}
```
