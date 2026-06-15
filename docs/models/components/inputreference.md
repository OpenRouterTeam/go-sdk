# InputReference

A reference asset used to guide video generation. Image references are supported by all providers; audio and video references are only honored by providers that support them (currently BytePlus Seedance 2.0).


## Supported Types

### ContentPartAudio

```go
inputReference := components.CreateInputReferenceAudioURL(components.ContentPartAudio{/* values here */})
```

### ContentPartImage

```go
inputReference := components.CreateInputReferenceImageURL(components.ContentPartImage{/* values here */})
```

### ContentPartVideo

```go
inputReference := components.CreateInputReferenceVideoURL(components.ContentPartVideo{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch inputReference.Type {
	case components.InputReferenceTypeAudioURL:
		// inputReference.ContentPartAudio is populated
	case components.InputReferenceTypeImageURL:
		// inputReference.ContentPartImage is populated
	case components.InputReferenceTypeVideoURL:
		// inputReference.ContentPartVideo is populated
}
```
