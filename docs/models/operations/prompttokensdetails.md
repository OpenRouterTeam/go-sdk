# PromptTokensDetails

Per-modality token breakdown. Only present when the input contains 2+ modalities (e.g. text + image) and the upstream provider returns modality-level usage data. Only non-zero modality counts are included.


## Fields

| Field                               | Type                                | Required                            | Description                         | Example                             |
| ----------------------------------- | ----------------------------------- | ----------------------------------- | ----------------------------------- | ----------------------------------- |
| `AudioTokens`                       | `*int64`                            | :heavy_minus_sign:                  | Number of audio tokens in the input |                                     |
| `ImageTokens`                       | `*int64`                            | :heavy_minus_sign:                  | Number of image tokens in the input | 258                                 |
| `TextTokens`                        | `*int64`                            | :heavy_minus_sign:                  | Number of text tokens in the input  | 8                                   |
| `VideoTokens`                       | `*int64`                            | :heavy_minus_sign:                  | Number of video tokens in the input |                                     |