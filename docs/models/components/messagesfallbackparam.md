# MessagesFallbackParam

Fallback model to try when the primary model fails or refuses. Only the `model` field is supported; per-attempt overrides are rejected.


## Fields

| Field                          | Type                           | Required                       | Description                    | Example                        |
| ------------------------------ | ------------------------------ | ------------------------------ | ------------------------------ | ------------------------------ |
| `Model`                        | `string`                       | :heavy_check_mark:             | N/A                            |                                |
| `AdditionalProperties`         | map[string]`any`               | :heavy_minus_sign:             | N/A                            | {<br/>"model": "claude-opus-4-8"<br/>} |