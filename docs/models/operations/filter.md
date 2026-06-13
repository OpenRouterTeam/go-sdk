# Filter


## Fields

| Field                                                  | Type                                                   | Required                                               | Description                                            | Example                                                |
| ------------------------------------------------------ | ------------------------------------------------------ | ------------------------------------------------------ | ------------------------------------------------------ | ------------------------------------------------------ |
| `Field`                                                | `string`                                               | :heavy_check_mark:                                     | Dimension to filter on                                 | model                                                  |
| `Operator`                                             | `string`                                               | :heavy_check_mark:                                     | Filter operator                                        | eq                                                     |
| `Value`                                                | [operations.Value1](../../models/operations/value1.md) | :heavy_check_mark:                                     | Filter value (scalar or array depending on operator)   |                                                        |