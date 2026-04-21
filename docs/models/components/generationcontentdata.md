# GenerationContentData

Stored prompt and completion content


## Fields

| Field                                                                                            | Type                                                                                             | Required                                                                                         | Description                                                                                      |
| ------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------ |
| `Input`                                                                                          | [components.InputUnion](../../models/components/inputunion.md)                                   | :heavy_check_mark:                                                                               | The input to the generation — either a prompt string or an array of messages                     |
| `Output`                                                                                         | [components.GenerationContentDataOutput](../../models/components/generationcontentdataoutput.md) | :heavy_check_mark:                                                                               | The output from the generation                                                                   |