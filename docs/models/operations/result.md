# Result

A single rerank result


## Fields

| Field                                                      | Type                                                       | Required                                                   | Description                                                | Example                                                    |
| ---------------------------------------------------------- | ---------------------------------------------------------- | ---------------------------------------------------------- | ---------------------------------------------------------- | ---------------------------------------------------------- |
| `Document`                                                 | [operations.Document](../../models/operations/document.md) | :heavy_check_mark:                                         | The document object containing the original text           |                                                            |
| `Index`                                                    | `int64`                                                    | :heavy_check_mark:                                         | Index of the document in the original input list           | 0                                                          |
| `RelevanceScore`                                           | `float64`                                                  | :heavy_check_mark:                                         | Relevance score of the document to the query               | 0.98                                                       |