# CapabilityDescriptor

A typed descriptor for one supported request parameter.


## Supported Types

### EnumCapability

```go
capabilityDescriptor := components.CreateCapabilityDescriptorEnum(components.EnumCapability{/* values here */})
```

### RangeCapability

```go
capabilityDescriptor := components.CreateCapabilityDescriptorRange(components.RangeCapability{/* values here */})
```

### BooleanCapability

```go
capabilityDescriptor := components.CreateCapabilityDescriptorBoolean(components.BooleanCapability{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch capabilityDescriptor.Type {
	case components.CapabilityDescriptorTypeEnum:
		// capabilityDescriptor.EnumCapability is populated
	case components.CapabilityDescriptorTypeRange:
		// capabilityDescriptor.RangeCapability is populated
	case components.CapabilityDescriptorTypeBoolean:
		// capabilityDescriptor.BooleanCapability is populated
	default:
		// Unknown type - use capabilityDescriptor.GetUnknownRaw() for raw JSON
}
```
