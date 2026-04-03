# ListOrganizationMembersData


## Fields

| Field                                              | Type                                               | Required                                           | Description                                        | Example                                            |
| -------------------------------------------------- | -------------------------------------------------- | -------------------------------------------------- | -------------------------------------------------- | -------------------------------------------------- |
| `ID`                                               | `string`                                           | :heavy_check_mark:                                 | User ID of the organization member                 | user_2dHFtVWx2n56w6HkM0000000000                   |
| `FirstName`                                        | `*string`                                          | :heavy_check_mark:                                 | First name of the member                           | Jane                                               |
| `LastName`                                         | `*string`                                          | :heavy_check_mark:                                 | Last name of the member                            | Doe                                                |
| `Email`                                            | `string`                                           | :heavy_check_mark:                                 | Email address of the member                        | jane.doe@example.com                               |
| `Role`                                             | [operations.Role](../../models/operations/role.md) | :heavy_check_mark:                                 | Role of the member in the organization             | org:member                                         |