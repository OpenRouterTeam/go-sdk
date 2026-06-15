# SubagentNestedTool

A tool made available to the subagent. Only OpenRouter server tools (e.g. openrouter:web_search) are supported; function tools are rejected because the worker has no way to execute them. The subagent tool may not list itself.


## Fields

| Field                               | Type                                | Required                            | Description                         | Example                             |
| ----------------------------------- | ----------------------------------- | ----------------------------------- | ----------------------------------- | ----------------------------------- |
| `Parameters`                        | map[string]`any`                    | :heavy_minus_sign:                  | N/A                                 |                                     |
| `Type`                              | `string`                            | :heavy_check_mark:                  | N/A                                 |                                     |
| `AdditionalProperties`              | map[string]`any`                    | :heavy_minus_sign:                  | N/A                                 | {<br/>"type": "openrouter:web_search"<br/>} |