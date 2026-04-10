# ListOrganizationMembersData


## Fields

| Field                                              | Type                                               | Required                                           | Description                                        | Example                                            |
| -------------------------------------------------- | -------------------------------------------------- | -------------------------------------------------- | -------------------------------------------------- | -------------------------------------------------- |
| `Email`                                            | `string`                                           | :heavy_check_mark:                                 | Email address of the member                        | jane.doe@example.com                               |
| `FirstName`                                        | `*string`                                          | :heavy_check_mark:                                 | First name of the member                           | Jane                                               |
| `ID`                                               | `string`                                           | :heavy_check_mark:                                 | User ID of the organization member                 | user_2dHFtVWx2n56w6HkM0000000000                   |
| `LastName`                                         | `*string`                                          | :heavy_check_mark:                                 | Last name of the member                            | Doe                                                |
| `Role`                                             | [operations.Role](../../models/operations/role.md) | :heavy_check_mark:                                 | Role of the member in the organization             | org:member                                         |