# CacheControl

Enable automatic prompt caching. When set, the system automatically applies cache breakpoints to the last cacheable block in the request. Currently supported for Anthropic Claude models.


## Fields

| Field                                                                    | Type                                                                     | Required                                                                 | Description                                                              |
| ------------------------------------------------------------------------ | ------------------------------------------------------------------------ | ------------------------------------------------------------------------ | ------------------------------------------------------------------------ |
| `Type`                                                                   | [components.ChatRequestType](../../models/components/chatrequesttype.md) | :heavy_check_mark:                                                       | N/A                                                                      |
| `TTL`                                                                    | [*components.ChatRequestTTL](../../models/components/chatrequestttl.md)  | :heavy_minus_sign:                                                       | N/A                                                                      |