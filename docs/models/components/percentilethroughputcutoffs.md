# PercentileThroughputCutoffs

Percentile-based throughput cutoffs. All specified cutoffs must be met for an endpoint to be preferred.


## Fields

| Field                                        | Type                                         | Required                                     | Description                                  |
| -------------------------------------------- | -------------------------------------------- | -------------------------------------------- | -------------------------------------------- |
| `P50`                                        | optionalnullable.OptionalNullable[`float64`] | :heavy_minus_sign:                           | Minimum p50 throughput (tokens/sec)          |
| `P75`                                        | optionalnullable.OptionalNullable[`float64`] | :heavy_minus_sign:                           | Minimum p75 throughput (tokens/sec)          |
| `P90`                                        | optionalnullable.OptionalNullable[`float64`] | :heavy_minus_sign:                           | Minimum p90 throughput (tokens/sec)          |
| `P99`                                        | optionalnullable.OptionalNullable[`float64`] | :heavy_minus_sign:                           | Minimum p99 throughput (tokens/sec)          |