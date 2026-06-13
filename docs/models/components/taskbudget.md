# TaskBudget

Task budget for an agentic turn. The model sees a countdown of remaining tokens and uses it to prioritize work and wind down gracefully. Advisory — does not enforce a hard cap.


## Fields

| Field                                                          | Type                                                           | Required                                                       | Description                                                    |
| -------------------------------------------------------------- | -------------------------------------------------------------- | -------------------------------------------------------------- | -------------------------------------------------------------- |
| `Remaining`                                                    | optionalnullable.OptionalNullable[`int64`]                     | :heavy_minus_sign:                                             | N/A                                                            |
| `Total`                                                        | `int64`                                                        | :heavy_check_mark:                                             | N/A                                                            |
| `Type`                                                         | [components.TypeTokens](../../models/components/typetokens.md) | :heavy_check_mark:                                             | N/A                                                            |