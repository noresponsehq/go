# PolicyRestoreRequest

Publish a captured configuration as a new version. Version checks and publication are atomic. Longer retention does not revive expired records, and restored policy does not undo customer enforcement.


## Fields

| Field                                               | Type                                                | Required                                            | Description                                         | Example                                             |
| --------------------------------------------------- | --------------------------------------------------- | --------------------------------------------------- | --------------------------------------------------- | --------------------------------------------------- |
| `Version`                                           | `int64`                                             | :heavy_check_mark:                                  | N/A                                                 |                                                     |
| `ExpectedPolicyVersion`                             | `int64`                                             | :heavy_check_mark:                                  | N/A                                                 |                                                     |
| `ChangeReason`                                      | `string`                                            | :heavy_check_mark:                                  | N/A                                                 | Adjust thresholds after reviewing retained evidence |