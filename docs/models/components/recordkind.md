# RecordKind

What sort of value a field holds, which decides how it is cleaned, compared and searched again. It comes from the value itself where its shape is certain (an email, an IP, a URL, an international phone number), else from the answer committed for the source's column, else from the answer stored for it, else text. A value of any kind is returned. Each kind's description says what it holds and the form value takes.

## Example Usage

```go
import (
	"github.com/noresponsehq/go/v12/models/components"
)

value := components.RecordKindEmail

// Open enum: custom values can be created with a direct type cast
custom := components.RecordKind("custom_value")
```


## Values

| Name                       | Value                      |
| -------------------------- | -------------------------- |
| `RecordKindEmail`          | email                      |
| `RecordKindPhone`          | phone                      |
| `RecordKindUsername`       | username                   |
| `RecordKindAccountID`      | account_id                 |
| `RecordKindURL`            | url                        |
| `RecordKindDomain`         | domain                     |
| `RecordKindImage`          | image                      |
| `RecordKindPersonName`     | person_name                |
| `RecordKindGender`         | gender                     |
| `RecordKindAge`            | age                        |
| `RecordKindLanguage`       | language                   |
| `RecordKindOrganisation`   | organisation               |
| `RecordKindJobTitle`       | job_title                  |
| `RecordKindCountry`        | country                    |
| `RecordKindRegion`         | region                     |
| `RecordKindCity`           | city                       |
| `RecordKindAddress`        | address                    |
| `RecordKindPostcode`       | postcode                   |
| `RecordKindCoordinates`    | coordinates                |
| `RecordKindTimezone`       | timezone                   |
| `RecordKindDate`           | date                       |
| `RecordKindDateTime`       | date_time                  |
| `RecordKindDuration`       | duration                   |
| `RecordKindIP`             | ip                         |
| `RecordKindMacAddress`     | mac_address                |
| `RecordKindDeviceID`       | device_id                  |
| `RecordKindUserAgent`      | user_agent                 |
| `RecordKindPassword`       | password                   |
| `RecordKindHash`           | hash                       |
| `RecordKindSalt`           | salt                       |
| `RecordKindToken`          | token                      |
| `RecordKindSecurityAnswer` | security_answer            |
| `RecordKindCardNumber`     | card_number                |
| `RecordKindBankAccount`    | bank_account               |
| `RecordKindCryptoAddress`  | crypto_address             |
| `RecordKindAmount`         | amount                     |
| `RecordKindIDNumber`       | id_number                  |
| `RecordKindVehicle`        | vehicle                    |
| `RecordKindLabel`          | label                      |
| `RecordKindLongText`       | long_text                  |
| `RecordKindCount`          | count                      |
| `RecordKindScore`          | score                      |
| `RecordKindFlag`           | flag                       |
| `RecordKindStructured`     | structured                 |
| `RecordKindText`           | text                       |