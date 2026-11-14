# RecordRole

Which common field a value is, where many sources share it: a join date and a birth date are both dates, with the roles joined and born. A platform account's field and a breach row's column with the same meaning carry the same kind and role. A linked_account field is named by its network. Absent for a field only its source returns. Each role's description names the kinds its values have.

## Example Usage

```go
import (
	"github.com/noresponsehq/go/v12/models/components"
)

value := components.RecordRoleAccountID

// Open enum: custom values can be created with a direct type cast
custom := components.RecordRole("custom_value")
```


## Values

| Name                        | Value                       |
| --------------------------- | --------------------------- |
| `RecordRoleAccountID`       | account_id                  |
| `RecordRoleAccountStatus`   | account_status              |
| `RecordRoleAccountType`     | account_type                |
| `RecordRoleAlternatePhone`  | alternate_phone             |
| `RecordRoleBalance`         | balance                     |
| `RecordRoleBanned`          | banned                      |
| `RecordRoleBio`             | bio                         |
| `RecordRoleBorn`            | born                        |
| `RecordRoleBot`             | bot                         |
| `RecordRoleCallerCategory`  | caller_category             |
| `RecordRoleChargebacks`     | chargebacks                 |
| `RecordRoleCoverPicture`    | cover_picture               |
| `RecordRoleDeclines`        | declines                    |
| `RecordRoleDeleted`         | deleted                     |
| `RecordRoleDeletedOn`       | deleted_on                  |
| `RecordRoleDepartment`      | department                  |
| `RecordRoleDisplayName`     | display_name                |
| `RecordRoleEmailVerified`   | email_verified              |
| `RecordRoleExpires`         | expires                     |
| `RecordRoleFirstName`       | first_name                  |
| `RecordRoleFollowers`       | followers                   |
| `RecordRoleFollowing`       | following                   |
| `RecordRoleFriends`         | friends                     |
| `RecordRoleFullAddress`     | full_address                |
| `RecordRoleFullName`        | full_name                   |
| `RecordRoleHandle`          | handle                      |
| `RecordRoleHeadline`        | headline                    |
| `RecordRoleInterests`       | interests                   |
| `RecordRoleJobTitle`        | job_title                   |
| `RecordRoleJoined`          | joined                      |
| `RecordRoleLastActive`      | last_active                 |
| `RecordRoleLastIP`          | last_ip                     |
| `RecordRoleLastName`        | last_name                   |
| `RecordRoleLastPost`        | last_post                   |
| `RecordRoleLastUpdated`     | last_updated                |
| `RecordRoleLevel`           | level                       |
| `RecordRoleLikesGiven`      | likes_given                 |
| `RecordRoleLikesReceived`   | likes_received              |
| `RecordRoleLineType`        | line_type                   |
| `RecordRoleLinkedAccount`   | linked_account              |
| `RecordRoleLocation`        | location                    |
| `RecordRoleMarkedSpam`      | marked_spam                 |
| `RecordRoleMiddleName`      | middle_name                 |
| `RecordRoleNationality`     | nationality                 |
| `RecordRoleOrganisation`    | organisation                |
| `RecordRoleOtherIP`         | other_ip                    |
| `RecordRolePaid`            | paid                        |
| `RecordRolePasswordChanged` | password_changed            |
| `RecordRolePaymentMethod`   | payment_method              |
| `RecordRolePicture`         | picture                     |
| `RecordRolePoints`          | points                      |
| `RecordRolePosts`           | posts                       |
| `RecordRolePrivate`         | private                     |
| `RecordRoleProfileLink`     | profile_link                |
| `RecordRoleProfileViews`    | profile_views               |
| `RecordRoleRating`          | rating                      |
| `RecordRoleReferrals`       | referrals                   |
| `RecordRoleReferrer`        | referrer                    |
| `RecordRoleReportCount`     | report_count                |
| `RecordRoleReputation`      | reputation                  |
| `RecordRoleReputationLabel` | reputation_label            |
| `RecordRoleResidence`       | residence                   |
| `RecordRoleSalutation`      | salutation                  |
| `RecordRoleSchool`          | school                      |
| `RecordRoleSignature`       | signature                   |
| `RecordRoleSignupIP`        | signup_ip                   |
| `RecordRoleSiteRole`        | site_role                   |
| `RecordRoleSpamScore`       | spam_score                  |
| `RecordRoleStreet`          | street                      |
| `RecordRoleTrades`          | trades                      |
| `RecordRoleVerified`        | verified                    |
| `RecordRoleWarnings`        | warnings                    |
| `RecordRoleWebsite`         | website                     |