# CreateRerankUsage

Usage statistics


## Fields

| Field                                            | Type                                             | Required                                         | Description                                      | Example                                          |
| ------------------------------------------------ | ------------------------------------------------ | ------------------------------------------------ | ------------------------------------------------ | ------------------------------------------------ |
| `TotalTokens`                                    | `*int64`                                         | :heavy_minus_sign:                               | Total number of tokens used                      | 150                                              |
| `SearchUnits`                                    | `*int64`                                         | :heavy_minus_sign:                               | Number of search units consumed (Cohere billing) | 1                                                |
| `Cost`                                           | `*float64`                                       | :heavy_minus_sign:                               | Cost of the request in credits                   | 0.001                                            |