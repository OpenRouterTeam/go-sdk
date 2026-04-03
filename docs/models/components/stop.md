# Stop

Stop sequences (up to 4)


## Supported Types

### 

```go
stop := components.CreateStopStr(string{/* values here */})
```

### 

```go
stop := components.CreateStopArrayOfStr([]string{/* values here */})
```

### 

```go
stop := components.CreateStopAny(any{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch stop.Type {
	case components.StopTypeStr:
		// stop.Str is populated
	case components.StopTypeArrayOfStr:
		// stop.ArrayOfStr is populated
	case components.StopTypeAny:
		// stop.Any is populated
}
```
