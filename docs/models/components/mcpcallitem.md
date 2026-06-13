# McpCallItem

An MCP tool call with its output or error


## Fields

| Field                                                                    | Type                                                                     | Required                                                                 | Description                                                              |
| ------------------------------------------------------------------------ | ------------------------------------------------------------------------ | ------------------------------------------------------------------------ | ------------------------------------------------------------------------ |
| `Arguments`                                                              | `string`                                                                 | :heavy_check_mark:                                                       | N/A                                                                      |
| `Error`                                                                  | optionalnullable.OptionalNullable[`string`]                              | :heavy_minus_sign:                                                       | N/A                                                                      |
| `ID`                                                                     | `string`                                                                 | :heavy_check_mark:                                                       | N/A                                                                      |
| `Name`                                                                   | `string`                                                                 | :heavy_check_mark:                                                       | N/A                                                                      |
| `Output`                                                                 | optionalnullable.OptionalNullable[`string`]                              | :heavy_minus_sign:                                                       | N/A                                                                      |
| `ServerLabel`                                                            | `string`                                                                 | :heavy_check_mark:                                                       | N/A                                                                      |
| `Type`                                                                   | [components.McpCallItemType](../../models/components/mcpcallitemtype.md) | :heavy_check_mark:                                                       | N/A                                                                      |