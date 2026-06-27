# ImageGenStreamErrorEventError

Provider error details


## Fields

| Field                                                      | Type                                                       | Required                                                   | Description                                                |
| ---------------------------------------------------------- | ---------------------------------------------------------- | ---------------------------------------------------------- | ---------------------------------------------------------- |
| `Code`                                                     | optionalnullable.OptionalNullable[`string`]                | :heavy_minus_sign:                                         | Provider error code, when supplied                         |
| `Message`                                                  | `string`                                                   | :heavy_check_mark:                                         | Provider error message                                     |
| `Param`                                                    | optionalnullable.OptionalNullable[`string`]                | :heavy_minus_sign:                                         | Request parameter associated with the error, when supplied |
| `Type`                                                     | optionalnullable.OptionalNullable[`string`]                | :heavy_minus_sign:                                         | Provider error type, when supplied                         |