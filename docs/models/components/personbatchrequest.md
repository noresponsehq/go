# PersonBatchRequest

Adds or updates 1 to 500 people. A row naming an existing person of the group (the same external_id, else the same email, else the same phone) updates it.


## Fields

| Field                                                                    | Type                                                                     | Required                                                                 | Description                                                              |
| ------------------------------------------------------------------------ | ------------------------------------------------------------------------ | ------------------------------------------------------------------------ | ------------------------------------------------------------------------ |
| `People`                                                                 | [][components.PersonIdentity](../../models/components/personidentity.md) | :heavy_check_mark:                                                       | N/A                                                                      |