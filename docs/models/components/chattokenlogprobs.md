# ChatTokenLogprobs

Log probabilities for the completion


## Fields

| Field                                                                                                           | Type                                                                                                            | Required                                                                                                        | Description                                                                                                     |
| --------------------------------------------------------------------------------------------------------------- | --------------------------------------------------------------------------------------------------------------- | --------------------------------------------------------------------------------------------------------------- | --------------------------------------------------------------------------------------------------------------- |
| `Content`                                                                                                       | [][components.ChatTokenLogprob](../../models/components/chattokenlogprob.md)                                    | :heavy_check_mark:                                                                                              | Log probabilities for content tokens                                                                            |
| `Refusal`                                                                                                       | optionalnullable.OptionalNullable[[][components.ChatTokenLogprob](../../models/components/chattokenlogprob.md)] | :heavy_minus_sign:                                                                                              | Log probabilities for refusal tokens                                                                            |