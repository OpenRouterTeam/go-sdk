# TopProviderInfo

Information about the top provider for this model


## Fields

| Field                                           | Type                                            | Required                                        | Description                                     | Example                                         |
| ----------------------------------------------- | ----------------------------------------------- | ----------------------------------------------- | ----------------------------------------------- | ----------------------------------------------- |
| `ContextLength`                                 | optionalnullable.OptionalNullable[`int64`]      | :heavy_minus_sign:                              | Context length from the top provider            | 8192                                            |
| `IsModerated`                                   | `bool`                                          | :heavy_check_mark:                              | Whether the top provider moderates content      | true                                            |
| `MaxCompletionTokens`                           | optionalnullable.OptionalNullable[`int64`]      | :heavy_minus_sign:                              | Maximum completion tokens from the top provider | 4096                                            |