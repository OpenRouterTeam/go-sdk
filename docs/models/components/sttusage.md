# STTUsage

Aggregated usage statistics for the request


## Fields

| Field                                          | Type                                           | Required                                       | Description                                    | Example                                        |
| ---------------------------------------------- | ---------------------------------------------- | ---------------------------------------------- | ---------------------------------------------- | ---------------------------------------------- |
| `Cost`                                         | `*float64`                                     | :heavy_minus_sign:                             | Total cost of the request in USD               | 0.000508                                       |
| `InputTokens`                                  | `*int64`                                       | :heavy_minus_sign:                             | Number of input tokens billed for this request | 83                                             |
| `OutputTokens`                                 | `*int64`                                       | :heavy_minus_sign:                             | Number of output tokens generated              | 30                                             |
| `Seconds`                                      | `*float64`                                     | :heavy_minus_sign:                             | Duration of the input audio in seconds         | 9.2                                            |
| `TotalTokens`                                  | `*int64`                                       | :heavy_minus_sign:                             | Total number of tokens used (input + output)   | 113                                            |