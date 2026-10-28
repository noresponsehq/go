# CredentialRequest

Associate the SHA-256 of a high-entropy key issued by your service. Send the hash instead of the key. Repeating a hash returns its original association, including its retirement state. A hash associated with another account returns a conflict.


## Fields

| Field                                                            | Type                                                             | Required                                                         | Description                                                      | Example                                                          |
| ---------------------------------------------------------------- | ---------------------------------------------------------------- | ---------------------------------------------------------------- | ---------------------------------------------------------------- | ---------------------------------------------------------------- |
| `Sha256`                                                         | `string`                                                         | :heavy_check_mark:                                               | N/A                                                              | 5da2158606a4cb396280c842b578872bb64b96c4eb80dc02e6fb4e0e864b77c6 |