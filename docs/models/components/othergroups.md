# OtherGroups

Where else in the environment the person appears, read live from those groups: deleting one removes it here.


## Fields

| Field                                                                               | Type                                                                                | Required                                                                            | Description                                                                         |
| ----------------------------------------------------------------------------------- | ----------------------------------------------------------------------------------- | ----------------------------------------------------------------------------------- | ----------------------------------------------------------------------------------- |
| `Total`                                                                             | `int64`                                                                             | :heavy_check_mark:                                                                  | How many other groups include this person, of which items lists the 20 most recent. |
| `Items`                                                                             | [][components.OtherGroup](../../models/components/othergroup.md)                    | :heavy_check_mark:                                                                  | This person's memberships in the 20 most recent other groups, newest group first.   |