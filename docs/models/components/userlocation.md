# UserLocation

Approximate user location for location-biased search results. Passed through to native providers that support it (e.g. Anthropic).


## Fields

| Field                                                                            | Type                                                                             | Required                                                                         | Description                                                                      |
| -------------------------------------------------------------------------------- | -------------------------------------------------------------------------------- | -------------------------------------------------------------------------------- | -------------------------------------------------------------------------------- |
| `City`                                                                           | optionalnullable.OptionalNullable[`string`]                                      | :heavy_minus_sign:                                                               | N/A                                                                              |
| `Country`                                                                        | optionalnullable.OptionalNullable[`string`]                                      | :heavy_minus_sign:                                                               | N/A                                                                              |
| `Region`                                                                         | optionalnullable.OptionalNullable[`string`]                                      | :heavy_minus_sign:                                                               | N/A                                                                              |
| `Timezone`                                                                       | optionalnullable.OptionalNullable[`string`]                                      | :heavy_minus_sign:                                                               | N/A                                                                              |
| `Type`                                                                           | [components.WebSearchPluginType](../../models/components/websearchplugintype.md) | :heavy_check_mark:                                                               | N/A                                                                              |