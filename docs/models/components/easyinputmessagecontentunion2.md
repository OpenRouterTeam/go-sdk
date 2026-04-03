# EasyInputMessageContentUnion2


## Supported Types

### 

```go
easyInputMessageContentUnion2 := components.CreateEasyInputMessageContentUnion2ArrayOfEasyInputMessageContentUnion1([]components.EasyInputMessageContentUnion1{/* values here */})
```

### 

```go
easyInputMessageContentUnion2 := components.CreateEasyInputMessageContentUnion2Str(string{/* values here */})
```

### 

```go
easyInputMessageContentUnion2 := components.CreateEasyInputMessageContentUnion2Any(any{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch easyInputMessageContentUnion2.Type {
	case components.EasyInputMessageContentUnion2TypeArrayOfEasyInputMessageContentUnion1:
		// easyInputMessageContentUnion2.ArrayOfEasyInputMessageContentUnion1 is populated
	case components.EasyInputMessageContentUnion2TypeStr:
		// easyInputMessageContentUnion2.Str is populated
	case components.EasyInputMessageContentUnion2TypeAny:
		// easyInputMessageContentUnion2.Any is populated
}
```
