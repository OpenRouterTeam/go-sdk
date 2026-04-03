# TextDeltaEventLogprob

Log probability information for a token


## Fields

| Field                                                                                        | Type                                                                                         | Required                                                                                     | Description                                                                                  |
| -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- |
| `Logprob`                                                                                    | `float64`                                                                                    | :heavy_check_mark:                                                                           | N/A                                                                                          |
| `Token`                                                                                      | `string`                                                                                     | :heavy_check_mark:                                                                           | N/A                                                                                          |
| `TopLogprobs`                                                                                | [][components.TextDeltaEventTopLogprob](../../models/components/textdeltaeventtoplogprob.md) | :heavy_minus_sign:                                                                           | N/A                                                                                          |
| `Bytes`                                                                                      | []`int64`                                                                                    | :heavy_minus_sign:                                                                           | N/A                                                                                          |