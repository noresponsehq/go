# AttendeeReport

Everything known about one attendee.


## Fields

| Field                                                                    | Type                                                                     | Required                                                                 | Description                                                              |
| ------------------------------------------------------------------------ | ------------------------------------------------------------------------ | ------------------------------------------------------------------------ | ------------------------------------------------------------------------ |
| `Attendee`                                                               | [components.Attendee](../../models/components/attendee.md)               | :heavy_check_mark:                                                       | An attendee and where their screening and review stand.                  |
| `Profiles`                                                               | [][components.Profile](../../models/components/profile.md)               | :heavy_check_mark:                                                       | The accounts the screening matched, with how each was matched.           |
| `Findings`                                                               | [][components.Finding](../../models/components/finding.md)               | :heavy_check_mark:                                                       | What the screening flagged, with its evidence.                           |
| `Reviews`                                                                | [][components.AttendeeReview](../../models/components/attendeereview.md) | :heavy_check_mark:                                                       | Every recorded decision, oldest first.                                   |