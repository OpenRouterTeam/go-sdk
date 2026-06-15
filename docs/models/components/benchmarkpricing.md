# BenchmarkPricing

OpenRouter pricing per token for this model. Null if pricing is unavailable.


## Fields

| Field                                        | Type                                         | Required                                     | Description                                  | Example                                      |
| -------------------------------------------- | -------------------------------------------- | -------------------------------------------- | -------------------------------------------- | -------------------------------------------- |
| `Completion`                                 | `string`                                     | :heavy_check_mark:                           | Cost per output token (USD, decimal string). | 0.000015                                     |
| `Prompt`                                     | `string`                                     | :heavy_check_mark:                           | Cost per input token (USD, decimal string).  | 0.000003                                     |