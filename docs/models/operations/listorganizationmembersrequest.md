# ListOrganizationMembersRequest


## Fields

| Field                                         | Type                                          | Required                                      | Description                                   | Example                                       |
| --------------------------------------------- | --------------------------------------------- | --------------------------------------------- | --------------------------------------------- | --------------------------------------------- |
| `Offset`                                      | optionalnullable.OptionalNullable[`int64`]    | :heavy_minus_sign:                            | Number of records to skip for pagination      | 0                                             |
| `Limit`                                       | `*int64`                                      | :heavy_minus_sign:                            | Maximum number of records to return (max 100) | 50                                            |