# FusionAnalysisResult

Structured analysis produced by the fusion judge model.


## Fields

| Field                                                                      | Type                                                                       | Required                                                                   | Description                                                                |
| -------------------------------------------------------------------------- | -------------------------------------------------------------------------- | -------------------------------------------------------------------------- | -------------------------------------------------------------------------- |
| `BlindSpots`                                                               | []`string`                                                                 | :heavy_check_mark:                                                         | N/A                                                                        |
| `Consensus`                                                                | []`string`                                                                 | :heavy_check_mark:                                                         | N/A                                                                        |
| `Contradictions`                                                           | [][components.Contradiction](../../models/components/contradiction.md)     | :heavy_check_mark:                                                         | N/A                                                                        |
| `PartialCoverage`                                                          | [][components.PartialCoverage](../../models/components/partialcoverage.md) | :heavy_check_mark:                                                         | N/A                                                                        |
| `UniqueInsights`                                                           | [][components.UniqueInsight](../../models/components/uniqueinsight.md)     | :heavy_check_mark:                                                         | N/A                                                                        |