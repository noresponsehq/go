# WebhookEndpointDisabled

One of the environment's endpoints stopped receiving events: every delivery attempt to it failed from failing_since until disabled_at, 72 hours or more. Its queued deliveries are dropped. Fix the receiver and verify it again with POST /v1/webhooks/{endpoint_id}/test (type endpoint.verification) to resume delivery.


## Fields

| Field                                                            | Type                                                             | Required                                                         | Description                                                      | Example                                                          |
| ---------------------------------------------------------------- | ---------------------------------------------------------------- | ---------------------------------------------------------------- | ---------------------------------------------------------------- | ---------------------------------------------------------------- |
| `EndpointID`                                                     | `string`                                                         | :heavy_check_mark:                                               | The endpoint that was disabled.                                  | whe_01JA2Z0A1B2C3D4E5F6G7H8J9K                                   |
| `FailingSince`                                                   | `string`                                                         | :heavy_check_mark:                                               | The first failed attempt of the unbroken failure streak.         | 2026-10-20T09:58:31.000Z                                         |
| `DisabledAt`                                                     | `string`                                                         | :heavy_check_mark:                                               | When the endpoint was disabled; equal to the event's created_at. | 2026-10-23T10:00:00.000Z                                         |