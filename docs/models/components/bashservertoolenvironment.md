# BashServerToolEnvironment

Execution environment for the bash server tool.


## Supported Types

### ContainerAutoEnvironment

```go
bashServerToolEnvironment := components.CreateBashServerToolEnvironmentContainerAuto(components.ContainerAutoEnvironment{/* values here */})
```

### ContainerReferenceEnvironment

```go
bashServerToolEnvironment := components.CreateBashServerToolEnvironmentContainerReference(components.ContainerReferenceEnvironment{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch bashServerToolEnvironment.Type {
	case components.BashServerToolEnvironmentTypeContainerAuto:
		// bashServerToolEnvironment.ContainerAutoEnvironment is populated
	case components.BashServerToolEnvironmentTypeContainerReference:
		// bashServerToolEnvironment.ContainerReferenceEnvironment is populated
}
```
