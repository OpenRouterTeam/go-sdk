# TextDoneEvent

Event emitted when text streaming is complete


## Fields

| Field                                                                        | Type                                                                         | Required                                                                     | Description                                                                  |
| ---------------------------------------------------------------------------- | ---------------------------------------------------------------------------- | ---------------------------------------------------------------------------- | ---------------------------------------------------------------------------- |
| `ContentIndex`                                                               | `int64`                                                                      | :heavy_check_mark:                                                           | N/A                                                                          |
| `ItemID`                                                                     | `string`                                                                     | :heavy_check_mark:                                                           | N/A                                                                          |
| `Logprobs`                                                                   | [][components.StreamLogprob](../../models/components/streamlogprob.md)       | :heavy_check_mark:                                                           | N/A                                                                          |
| `OutputIndex`                                                                | `int64`                                                                      | :heavy_check_mark:                                                           | N/A                                                                          |
| `SequenceNumber`                                                             | `int64`                                                                      | :heavy_check_mark:                                                           | N/A                                                                          |
| `Text`                                                                       | `string`                                                                     | :heavy_check_mark:                                                           | N/A                                                                          |
| `Type`                                                                       | [components.TextDoneEventType](../../models/components/textdoneeventtype.md) | :heavy_check_mark:                                                           | N/A                                                                          |