# AttendeeReport

Everything known about one attendee.


## Fields

| Field                                                                           | Type                                                                            | Required                                                                        | Description                                                                     |
| ------------------------------------------------------------------------------- | ------------------------------------------------------------------------------- | ------------------------------------------------------------------------------- | ------------------------------------------------------------------------------- |
| `Attendee`                                                                      | [components.Attendee](../../models/components/attendee.md)                      | :heavy_check_mark:                                                              | An attendee and where their screening and review stand.                         |
| `Profiles`                                                                      | [][components.ScreeningProfile](../../models/components/screeningprofile.md)    | :heavy_check_mark:                                                              | The accounts the latest screening matched, with how each was matched.           |
| `Findings`                                                                      | [][components.ScreeningFinding](../../models/components/screeningfinding.md)    | :heavy_check_mark:                                                              | What the latest screening flagged, most severe first, each citing its evidence. |
| `Posts`                                                                         | [][components.ScreeningPost](../../models/components/screeningpost.md)          | :heavy_check_mark:                                                              | The posts the findings cite.                                                    |
| `Media`                                                                         | [][components.ScreeningMedia](../../models/components/screeningmedia.md)        | :heavy_check_mark:                                                              | The images and videos the findings cite.                                        |
| `Reviews`                                                                       | [][components.AttendeeReview](../../models/components/attendeereview.md)        | :heavy_check_mark:                                                              | Every recorded decision, oldest first.                                          |