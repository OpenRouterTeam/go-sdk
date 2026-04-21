# SpeechRequestProvider

Provider-specific passthrough configuration


## Fields

| Field                                                                                                                             | Type                                                                                                                              | Required                                                                                                                          | Description                                                                                                                       |
| --------------------------------------------------------------------------------------------------------------------------------- | --------------------------------------------------------------------------------------------------------------------------------- | --------------------------------------------------------------------------------------------------------------------------------- | --------------------------------------------------------------------------------------------------------------------------------- |
| `Options`                                                                                                                         | [*components.SpeechRequestOptions](../../models/components/speechrequestoptions.md)                                               | :heavy_minus_sign:                                                                                                                | Provider-specific options keyed by provider slug. The options for the matched provider are spread into the upstream request body. |