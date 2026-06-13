# LocalShellCallItemAction


## Fields

| Field                                                      | Type                                                       | Required                                                   | Description                                                |
| ---------------------------------------------------------- | ---------------------------------------------------------- | ---------------------------------------------------------- | ---------------------------------------------------------- |
| `Command`                                                  | []`string`                                                 | :heavy_check_mark:                                         | N/A                                                        |
| `Env`                                                      | map[string]`string`                                        | :heavy_check_mark:                                         | N/A                                                        |
| `TimeoutMs`                                                | optionalnullable.OptionalNullable[`int64`]                 | :heavy_minus_sign:                                         | N/A                                                        |
| `Type`                                                     | [components.TypeExec](../../models/components/typeexec.md) | :heavy_check_mark:                                         | N/A                                                        |
| `User`                                                     | optionalnullable.OptionalNullable[`string`]                | :heavy_minus_sign:                                         | N/A                                                        |
| `WorkingDirectory`                                         | optionalnullable.OptionalNullable[`string`]                | :heavy_minus_sign:                                         | N/A                                                        |