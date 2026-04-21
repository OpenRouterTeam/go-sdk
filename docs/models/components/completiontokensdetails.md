# CompletionTokensDetails

Detailed completion token usage


## Fields

| Field                                      | Type                                       | Required                                   | Description                                |
| ------------------------------------------ | ------------------------------------------ | ------------------------------------------ | ------------------------------------------ |
| `AcceptedPredictionTokens`                 | optionalnullable.OptionalNullable[`int64`] | :heavy_minus_sign:                         | Accepted prediction tokens                 |
| `AudioTokens`                              | optionalnullable.OptionalNullable[`int64`] | :heavy_minus_sign:                         | Tokens used for audio output               |
| `ReasoningTokens`                          | optionalnullable.OptionalNullable[`int64`] | :heavy_minus_sign:                         | Tokens used for reasoning                  |
| `RejectedPredictionTokens`                 | optionalnullable.OptionalNullable[`int64`] | :heavy_minus_sign:                         | Rejected prediction tokens                 |