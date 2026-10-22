# AssessmentRequestIssuedCredential

Optional. When the signup gate issues an API key to the new account, send SHA-256 of that key, so later usage of the key, or a leaked copy, ties back to this assessment. Each SDK has a helper that computes it. Send only the hash: customer keys are high-entropy, so the hash reveals nothing about them. The assessment stores it, and no response, webhook or log carries it.


## Fields

| Field                                                                     | Type                                                                      | Required                                                                  | Description                                                               | Example                                                                   |
| ------------------------------------------------------------------------- | ------------------------------------------------------------------------- | ------------------------------------------------------------------------- | ------------------------------------------------------------------------- | ------------------------------------------------------------------------- |
| `Sha256`                                                                  | `string`                                                                  | :heavy_check_mark:                                                        | SHA-256 of the issued key's UTF-8 bytes, as 64 lower-case hex characters. | 2bdbf5c72f2a36df0ef40fb92203612274d4546f56dba946070a6a79c2401a1b          |