# PreferredMaxLatency

Preferred maximum latency (in seconds). Can be a number (applies to p50) or an object with percentile-specific cutoffs. Endpoints above the threshold(s) may still be used, but are deprioritized in routing. When using fallback models, this may cause a fallback model to be used instead of the primary model if it meets the threshold.


## Supported Types

### 

```go
preferredMaxLatency := components.CreatePreferredMaxLatencyNumber(float64{/* values here */})
```

### PercentileLatencyCutoffs

```go
preferredMaxLatency := components.CreatePreferredMaxLatencyPercentileLatencyCutoffs(components.PercentileLatencyCutoffs{/* values here */})
```

### 

```go
preferredMaxLatency := components.CreatePreferredMaxLatencyAny(any{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch preferredMaxLatency.Type {
	case components.PreferredMaxLatencyTypeNumber:
		// preferredMaxLatency.Number is populated
	case components.PreferredMaxLatencyTypePercentileLatencyCutoffs:
		// preferredMaxLatency.PercentileLatencyCutoffs is populated
	case components.PreferredMaxLatencyTypeAny:
		// preferredMaxLatency.Any is populated
}
```
