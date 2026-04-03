# CompletionTokensDetails

Detailed completion token usage


## Fields

| Field                        | Type                         | Required                     | Description                  |
| ---------------------------- | ---------------------------- | ---------------------------- | ---------------------------- |
| `ReasoningTokens`            | `*int64`                     | :heavy_minus_sign:           | Tokens used for reasoning    |
| `AudioTokens`                | `*int64`                     | :heavy_minus_sign:           | Tokens used for audio output |
| `AcceptedPredictionTokens`   | `*int64`                     | :heavy_minus_sign:           | Accepted prediction tokens   |
| `RejectedPredictionTokens`   | `*int64`                     | :heavy_minus_sign:           | Rejected prediction tokens   |