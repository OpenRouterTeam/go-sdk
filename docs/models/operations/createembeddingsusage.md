# CreateEmbeddingsUsage

Token usage statistics


## Fields

| Field                          | Type                           | Required                       | Description                    | Example                        |
| ------------------------------ | ------------------------------ | ------------------------------ | ------------------------------ | ------------------------------ |
| `Cost`                         | `*float64`                     | :heavy_minus_sign:             | Cost of the request in credits | 0.0001                         |
| `PromptTokens`                 | `int64`                        | :heavy_check_mark:             | Number of tokens in the input  | 8                              |
| `TotalTokens`                  | `int64`                        | :heavy_check_mark:             | Total number of tokens used    | 8                              |