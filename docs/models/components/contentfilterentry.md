# ContentFilterEntry

A custom regex content filter that scans request messages for matching patterns.


## Fields

| Field                                                                            | Type                                                                             | Required                                                                         | Description                                                                      | Example                                                                          |
| -------------------------------------------------------------------------------- | -------------------------------------------------------------------------------- | -------------------------------------------------------------------------------- | -------------------------------------------------------------------------------- | -------------------------------------------------------------------------------- |
| `Action`                                                                         | [components.ContentFilterAction](../../models/components/contentfilteraction.md) | :heavy_check_mark:                                                               | Action taken when the pattern matches                                            | block                                                                            |
| `Label`                                                                          | optionalnullable.OptionalNullable[`string`]                                      | :heavy_minus_sign:                                                               | Optional label used in redaction placeholders or error messages                  | [API_KEY]                                                                        |
| `Pattern`                                                                        | `string`                                                                         | :heavy_check_mark:                                                               | A regex pattern to match against request content                                 | \b(sk-[a-zA-Z0-9]{48})\b                                                         |